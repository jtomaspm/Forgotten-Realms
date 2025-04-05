package controllers

import (
	"backend/lib/auth_server/configuration"
	"backend/lib/auth_server/dal/models/queries"
	"backend/lib/auth_server/dal/services/account"
	"backend/lib/auth_server/dal/services/login"
	"backend/lib/auth_server/dal/services/session"
	"backend/lib/auth_server/server/models"
	"backend/pkg/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type GithubController struct {
	Configuration *configuration.Configuration
	Database      *database.Database
}

func (controller *GithubController) Mount(basePath string, engine *gin.Engine) {
	path := basePath + "/github"
	engine.GET(path+"/login", controller.login)
	engine.GET(path+"/callback", controller.callback)
}

func (controller *GithubController) login(ctx *gin.Context) {
	scope := "user:email"
	url := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=%s",
		controller.Configuration.GitHub.ClientId,
		controller.Configuration.GitHub.RedirectUri,
		scope)
	ctx.Redirect(http.StatusFound, url)
}

func (controller *GithubController) callback(ctx *gin.Context) {
	//Valiadate code
	errorStatus := http.StatusBadRequest
	errorMsg := "Missing code parameter"
	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}

	//Get access token
	errorStatus = http.StatusUnauthorized
	errorMsg = "Failed to get access token"
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", nil)
	if err != nil {
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}
	urlQuery := req.URL.Query()
	urlQuery.Add("client_id", controller.Configuration.GitHub.ClientId)
	urlQuery.Add("client_secret", controller.Configuration.GitHub.ClientSecret)
	urlQuery.Add("code", code)
	req.URL.RawQuery = urlQuery.Encode()
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}
	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	resp.Body.Close()
	if err != nil {
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}
	accessToken := result["access_token"]

	//Get user from github
	errorMsg = "Failed to fetch user data"
	req, err = http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("User-Agent", controller.Configuration.Server.UserAgent)
	resp, err = client.Do(req)
	if err != nil {
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}
	var user struct {
		Id int `json:"id"`
	}
	if json.NewDecoder(resp.Body).Decode(&user) != nil {
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}
	resp.Body.Close()
	externalId := fmt.Sprintf("%d", user.Id)

	//Try to login user
	account, err := account.GetByExternalId(ctx, controller.Database.Pool, &queries.GetAccountByExternalId{
		ExternalId: externalId,
		Source:     controller.Configuration.GitHub.Source,
	})
	if err == nil {
		_, err = login.Create(ctx, controller.Database.Pool, &queries.CreateLogin{
			AccountId: account.Id,
			IpAddress: ctx.ClientIP(),
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new login"})
			return
		}
		token, expires_at, err := session.Create(ctx, controller.Database.Pool, account.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new session"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"token":      token,
			"expires_at": expires_at,
		})
		return
	}

	//Get user email to create new account
	errorMsg = "Failed to fetch user email"
	req, err = http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("User-Agent", controller.Configuration.Server.UserAgent)
	resp, err = client.Do(req)
	if err != nil {
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}
	var emails []struct {
		Email    string `json:"email"`
		Primary  bool   `json:"primary"`
		Verified bool   `json:"verified"`
	}
	if json.NewDecoder(resp.Body).Decode(&emails) != nil {
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}
	resp.Body.Close()
	var email string
	for _, e := range emails {
		if e.Primary && e.Verified {
			email = e.Email
			break
		}
	}
	if email == "" {
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}

	//Provide needed fields new account flow
	errorStatus = http.StatusInternalServerError
	errorMsg = "Failed to generate new user claims"
	expires_at := time.Now().Add(10 * time.Minute)
	claims := models.NewUserClaims{
		ExternalId: externalId,
		Source:     controller.Configuration.GitHub.Source,
		Email:      email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    controller.Configuration.Server.UserAgent,
			ExpiresAt: jwt.NewNumericDate(expires_at),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(controller.Configuration.JwtSecret))
	if err != nil {
		log.Println(err)
		ctx.JSON(errorStatus, gin.H{"error": errorMsg})
		return
	}
	ctx.JSON(http.StatusMethodNotAllowed, gin.H{
		"token":      signedToken,
		"email":      email,
		"expires_at": expires_at,
		"message":    "Create new account using token",
	})
}

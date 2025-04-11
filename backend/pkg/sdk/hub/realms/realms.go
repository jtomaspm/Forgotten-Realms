package realms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type RegisterRealmRequest struct {
	Hub           string
	InternalToken string
	Body          *RegisterRealmRequestBody
}

type RegisterRealmRequestBody struct {
	Name string `json:"name"`
	Api  string `json:"api"`
}

func RegisterRealm(request *RegisterRealmRequest) (uuid.UUID, error) {
	var response struct {
		Id uuid.UUID `json:"id"`
	}

	reqPath := "http://" + request.Hub + "/api/realm"

	bodyBytes, err := json.Marshal(request.Body)
	if err != nil {
		return response.Id, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", reqPath, bytes.NewReader(bodyBytes))
	if err != nil {
		return response.Id, err
	}

	req.Header.Set("Authorization", "Internal "+request.InternalToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response.Id, err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
		return response.Id, fmt.Errorf("failed to register realm: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Id, err
	}
	return response.Id, nil
}

type RegisterAccountRequest struct {
	Hub           string
	InternalToken string
	Body          *RegisterAccountRequestBody
}

type RegisterAccountRequestBody struct {
	AccountId uuid.UUID `json:"account_id"`
	RealmId   uuid.UUID `json:"realm_id"`
}

func RegisterAccount(command *RegisterAccountRequest) error {
	reqPath := "http://" + command.Hub + "/api/realm/account"

	bodyBytes, err := json.Marshal(command.Body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", reqPath, bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Internal "+command.InternalToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
		return fmt.Errorf("failed to register account: %s", resp.Status)
	}
	return nil
}

type GetRealmRequest struct {
	Hub     string
	RealmId uuid.UUID
}

func IsRealmOpen(command *GetRealmRequest) (bool, error) {
	reqPath := "http://" + command.Hub + "/api/realm/" + command.RealmId.String()

	req, err := http.NewRequest("GET", reqPath, nil)
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode == http.StatusOK) {
		return false, fmt.Errorf("failed to register account: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	log.Println(string(body))
	if err != nil {
		return false, err
	}
	var realm struct {
		Status string `json:"status"`
	}
	err = json.Unmarshal(body, &realm)
	if err != nil {
		return false, err
	}
	log.Println(realm, realm.Status)
	return realm.Status == "open", nil
}

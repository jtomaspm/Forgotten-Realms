using System.Net.Http.Headers;
using System.Text.Json;
using Database;
using Database.Application;
using Database.Application.Extensions;
using Database.Application.Models;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Mvc;

namespace Backend.Controllers.Auth
{
    [Route("api/auth/github")]
    [ApiController]
    public class GitHubAuthController(IConfiguration configuration, ILogger<GitHubAuthController> logger, IDatabaseFactory<ApplicationDatabase> databaseFactory) : ControllerBase
    {
        [HttpGet("test")]
        public async Task<IActionResult> GitHubUserInfo()
        {
            await HttpContext.GetTokenAsync("acess_token");

            var externalId = HttpContext.User.Claims.FirstOrDefault(x=>x.Type == "externalId");
            var email = HttpContext.User.Claims.FirstOrDefault(x=>x.Type == "email");
            if (email is not null && externalId is not null)
                logger.LogInformation($"logged in [ ExternalId: {externalId.Value.AsSpan().ToString()}, "
                                                + $"Email: {email.Value.AsSpan().ToString()} ]");
            return Ok(
                HttpContext.User.Claims
                    .Select(c => new{c.Type, c.Value})
                    .ToList()
            );
        }

        [HttpGet("login")]
        public IActionResult GitHubLogin()
        {
            string client_id = configuration["GITHUB_CLIENT_ID"]!;
            string redirect_uri = $"http://localhost:{configuration["BACKEND_HTTP_PORT"]}/api/auth/github/callback";
            string scope = "user:email";
            return Redirect($"https://github.com/login/oauth/authorize?client_id={client_id}&redirect_uri={redirect_uri}&scope={scope}");
        }


        [HttpGet("callback")]
        public async Task<IActionResult> Callback(string code)
        {
            logger.LogInformation("Starting callback...");
            var clientId = configuration["GITHUB_CLIENT_ID"]!;
            logger.LogInformation($"ClientId: {clientId}");
            var clientSecret = configuration["GITHUB_CLIENT_SECRET"]!;
            logger.LogInformation($"ClientSecret: {clientSecret}");

            var httpClient = new HttpClient();
            var tokenRequest = new HttpRequestMessage(HttpMethod.Post, "https://github.com/login/oauth/access_token")
            {
                Content = new FormUrlEncodedContent(
                [
                    new KeyValuePair<string, string>("client_id", clientId),
                    new KeyValuePair<string, string>("client_secret", clientSecret),
                    new KeyValuePair<string, string>("code", code),
                ])
            };

            tokenRequest.Headers.Accept.Add(new MediaTypeWithQualityHeaderValue("application/json"));

            var response = await httpClient.SendAsync(tokenRequest);
            var responseContent = await response.Content.ReadAsStringAsync();

            logger.LogInformation(responseContent);

            var tokenJson = JsonDocument.Parse(responseContent);
            var accessToken = tokenJson.RootElement.GetProperty("access_token").GetString();

            logger.LogInformation($"AccessToken: {accessToken}");

            if (string.IsNullOrEmpty(accessToken))
            {
                return Unauthorized("Failed to retrieve access token from GitHub.");
            }

            var userRequest = new HttpRequestMessage(HttpMethod.Get, "https://api.github.com/user");
            userRequest.Headers.Authorization = new AuthenticationHeaderValue("Bearer", accessToken);
            userRequest.Headers.UserAgent.ParseAdd("SimplifiedCrafter/Backend");

            var userResponse = await httpClient.SendAsync(userRequest);
            logger.LogInformation(await userResponse.Content.ReadAsStringAsync());

            var emailRequest = new HttpRequestMessage(HttpMethod.Get, "https://api.github.com/user/emails");
            emailRequest.Headers.Authorization = new AuthenticationHeaderValue("Bearer", accessToken);
            emailRequest.Headers.UserAgent.ParseAdd("SimplifiedCrafter/Backend");

            var externalId = (await userResponse.Content.ReadFromJsonAsync<JsonElement>())
                .GetProperty("id")
                .GetRawText();
            var source = "GitHub";

            logger.LogInformation($"ExternalId: {externalId}");
            if (externalId is null) 
                return Unauthorized("Failed to fetch user information from GitHub.");

            //TODO: Process the user data here, you can create a session or JWT, etc.
            using var database = await databaseFactory.CreateDatabase();
            var account = await database.ExecuteInTransaction(async _ => {
                var acc = await database.GetAccountByExternalId(externalId, source);

                logger.LogInformation("{@Account}", acc);
                AccountDetails result;
                if (acc is not null)
                {
                    result = new() 
                    {
                        Id = acc.Id,
                        ExternalId = acc.ExternalId,
                        Source = acc.Source,
                        Name = acc.Name,
                        Email = acc.Email,
                        Role = acc.Role,
                        CreatedAt = acc.CreatedAt,
                        UpdatedAt = acc.UpdatedAt,
                        AccountProperties = await database.GetAccountPropertiesById(acc.Id) ?? throw new Exception($"Account created without properties: {acc.Id}"),
                        LastLogin = null,
                        Session = null,
                        Worlds = [], //TODO
                    };
                }
                else
                {
                    var emailResponse = await httpClient.SendAsync(emailRequest);
                    var email = (await emailResponse.Content.ReadFromJsonAsync<JsonElement>())
                        .EnumerateArray()
                        .FirstOrDefault(e => e.GetProperty("primary").GetBoolean())
                        .GetProperty("email")
                        .GetString();
                    logger.LogInformation($"Email: {email}");

                    if (email is null)
                        return null;

                    result = await database.CreateAccount(source, externalId, null, email, Role.Player());
                }

                var ipAddress = HttpContext.Connection.RemoteIpAddress is not null ? HttpContext.Connection.RemoteIpAddress.ToString() : "";
                result.LastLogin = await database.CreateLogin(result.Id, ipAddress);
                result.Session = await database.CreateSession(result.Id);

                return result;
            });
            if (account is null) 
                return Unauthorized(new {Message="Error fetching user data from GitHub."});
            return Ok(new 
            {
                account.Id, 
                account.ExternalId, 
                account.Source, 
                account.Email, 
                account.Name, 
                Role=account.Role.Name.ToString(), 
                account.CreatedAt, 
                account.UpdatedAt,
                Properties=new 
                {
                    account.AccountProperties.EmailVerified,
                    account.AccountProperties.SendEmailNotifications,
                    account.AccountProperties.VerificationToken,
                    account.AccountProperties.TokenExpiresAt,
                },
                LastLogin=new{
                    account.LastLogin!.IpAddress,
                    account.LastLogin!.CreatedAt
                },
                Session=new{
                    Sid=account.Session!.Id,
                    account.Session!.Token,
                    account.Session!.ExpiresAt,
                }
            });

        }

        [HttpGet("logout")]
        public async Task<IActionResult> Logout()
        {
            await HttpContext.SignOutAsync("GitHub");
            return Ok(new { Message = "Logged out successfully" });
        }
    }
}

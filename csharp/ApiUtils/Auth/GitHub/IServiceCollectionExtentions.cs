using Microsoft.Extensions.DependencyInjection;
using Microsoft.AspNetCore.Authentication;
using System.Security.Claims;
using System.Text.Json;
using Microsoft.AspNetCore.Authentication.Cookies;


namespace ApiUtils.Auth.GitHub;

public static class IServiceCollectionExtentions
{
    public static void SetupGithubAuth(this IServiceCollection services, string clientId, string clientSecret, string redirectUri)
    {
        services
            .AddAuthentication(options =>
            {
                options.DefaultScheme = CookieAuthenticationDefaults.AuthenticationScheme;
                options.DefaultChallengeScheme = "GitHub";
            })
            .AddCookie()
            .AddOAuth("GitHub", options =>
            {
                options.ClientId = clientId;
                options.ClientSecret = clientSecret;
                options.CallbackPath = redirectUri;

                options.AuthorizationEndpoint = "https://github.com/login/oauth/authorize";
                options.TokenEndpoint = "https://github.com/login/oauth/access_token";
                options.UserInformationEndpoint = "https://api.github.com/user";

                options.ClaimActions.MapJsonKey(ClaimTypes.NameIdentifier, "id");
                options.ClaimActions.MapJsonKey(ClaimTypes.Name, "login");
                options.ClaimActions.MapJsonKey(ClaimTypes.Email, "email");

                options.Events.OnCreatingTicket = async context =>
                {
                    var userInfo = await context.Backchannel.GetAsync(context.Options.UserInformationEndpoint);
                    if (userInfo.IsSuccessStatusCode)
                    {
                        using var jsonDoc = JsonDocument.Parse(await userInfo.Content.ReadAsStringAsync());
                        var root = jsonDoc.RootElement;

                        var githubId = root.GetProperty("id").GetInt64();
                        var username = root.GetProperty("login").GetString();
                        var email = root.TryGetProperty("email", out var emailProp) ? emailProp.GetString() : null;

                        context.Identity?.AddClaim(new Claim("GitHubId", githubId.ToString()));
                        context.Identity?.AddClaim(new Claim(ClaimTypes.Name, username));
                        context.Identity?.AddClaim(new Claim(ClaimTypes.Email, email ?? ""));
                    }
                };
            });
    }
}

using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Security.Claims;
using System.Text.Json;
using Microsoft.AspNetCore.Authentication;
using Microsoft.Extensions.DependencyInjection;


namespace ApiUtils.Auth.GitHub;

public static class IServiceCollectionExtentions
{
    public static void SetupGithubAuth(this IServiceCollection services, string clientId, string clientSecret, string redirectUri)
    {
        services
            .AddAuthentication("GitHubCookie")
            .AddCookie("GitHubCookie")
            .AddOAuth("GitHub", options =>
            {
                options.SignInScheme = "GitHubCookie";
                options.SaveTokens = true;

                options.ClientId = clientId;
                options.ClientSecret = clientSecret;
                options.CallbackPath = redirectUri;

                options.AuthorizationEndpoint = "https://github.com/login/oauth/authorize";
                options.TokenEndpoint = "https://github.com/login/oauth/access_token";
                options.UserInformationEndpoint = "https://api.github.com/user";

                options.Scope.Add("user:email");

                options.ClaimActions.MapJsonKey("externalId", "id");
                options.ClaimActions.MapJsonKey("email", "email");

                options.Events.OnCreatingTicket = async (ctx) => 
                {
                    var request = new HttpRequestMessage(HttpMethod.Get, ctx.Options.UserInformationEndpoint);
                    request.Headers.Authorization = new AuthenticationHeaderValue("Bearer", ctx.AccessToken);
                    using var userResponse = await ctx.Backchannel.SendAsync(request);
                    var user = await userResponse.Content.ReadFromJsonAsync<JsonElement>();
                    ctx.RunClaimActions(user);

                    // If the email is missing (private email), fetch from /user/emails
                    if (ctx.Identity is not null && !ctx.Identity.HasClaim(c => c.Type == "email"))
                    {
                        var emailRequest = new HttpRequestMessage(HttpMethod.Get, "https://api.github.com/user/emails");
                        emailRequest.Headers.Authorization = new AuthenticationHeaderValue("Bearer", ctx.AccessToken);
                        emailRequest.Headers.Accept.Add(new MediaTypeWithQualityHeaderValue("application/json"));

                        using var emailResponse = await ctx.Backchannel.SendAsync(emailRequest);
                        if (emailResponse.IsSuccessStatusCode)
                        {
                            var emails = await emailResponse.Content.ReadFromJsonAsync<JsonElement>();
                            var primaryEmail = emails.EnumerateArray()
                                .FirstOrDefault(e => e.GetProperty("primary").GetBoolean())
                                .GetProperty("email").GetString();

                            if (!string.IsNullOrEmpty(primaryEmail))
                            {
                                ctx.Identity?.AddClaim(new Claim("email", primaryEmail));
                            }
                        }
                    }
                };
            });
    }
}

using System.Security.Claims;
using System.Text.Json;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.Cookies;
using MySql.Data.MySqlClient;

namespace Backend;

public class Startup
{
    private readonly IConfiguration _configuration;
    public Startup(IConfiguration configuration)
    {
        _configuration = configuration ?? throw new ArgumentNullException(nameof(configuration));
        
        if (_configuration["GITHUB_CLIENT_ID"] is null) throw new ArgumentException("GITHUB_CLIENT_ID environment variable is not defined.");
        if (_configuration["GITHUB_CLIENT_SECRET"] is null) throw new ArgumentException("GITHUB_CLIENT_SECRET environment variable is not defined.");
        
        if (_configuration["MYSQL_DATABASE"] is null) throw new ArgumentException("MYSQL_DATABASE environment variable is not defined.");
        if (_configuration["MYSQL_PORT"] is null) throw new ArgumentException("MYSQL_PORT environment variable is not defined.");
        if (_configuration["MYSQL_HOST"] is null) throw new ArgumentException("MYSQL_HOST environment variable is not defined.");
        if (_configuration["MYSQL_ROOT_PASSWORD"] is null) throw new ArgumentException("MYSQL_ROOT_PASSWORD environment variable is not defined.");
        
        _configuration["MYSQL_CONNECTION_STRING"] = $"Server={_configuration["MYSQL_HOST"]};Port={_configuration["MYSQL_PORT"]};Database={_configuration["MYSQL_DATABASE"]};User=root;Password={_configuration["MYSQL_ROOT_PASSWORD"]};";
    }
    public void ConfigureServices(IServiceCollection services)
    {
        services.AddControllers();
        
        // Authentication Middleware
        services
            .AddAuthentication(options =>
            {
                options.DefaultScheme = CookieAuthenticationDefaults.AuthenticationScheme;
                options.DefaultChallengeScheme = "GitHub";
            })
            .AddCookie()
            .AddOAuth("GitHub", options =>
            {
                options.ClientId = _configuration["GITHUB_CLIENT_ID"]!;
                options.ClientSecret = _configuration["GITHUB_CLIENT_SECRET"]!;
                options.CallbackPath= "/api/auth/github/callback";

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

                        var connectionString = _configuration["MYSQL_CONNECTION_STRING"];
                        using var connection = new MySqlConnection(connectionString);
                        await connection.OpenAsync();

                        var query = @"
                        INSERT INTO users (github_id, username, email) 
                        VALUES (@github_id, @username, @email)
                        ON DUPLICATE KEY UPDATE username = @username, email = @email";

                        using var command = new MySqlCommand(query, connection);
                        command.Parameters.AddWithValue("@github_id", githubId);
                        command.Parameters.AddWithValue("@username", username);
                        command.Parameters.AddWithValue("@email", email ?? (object)DBNull.Value);
                        await command.ExecuteNonQueryAsync();
                    }
                };
            });
    }

    public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
    {
        if (env.IsDevelopment())
        {
            app.UseDeveloperExceptionPage();
        }

        app.UseRouting();
        app.UseAuthentication();
        app.UseAuthorization();
        app.UseEndpoints(endpoints =>
        {
            endpoints.MapControllers();
        });
    }
}
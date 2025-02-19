using ApiUtils.Auth.GitHub;
using ApiUtils.Configuration;
using ApiUtils.Middleware;
using Database;

namespace Backend;

public class Startup
{
    private readonly IConfiguration _configuration;
    private readonly DatabaseConfig _databaseConfig;
    public Startup(IConfiguration configuration)
    {
        _configuration = configuration ?? throw new ArgumentNullException(nameof(configuration));
        _configuration.ValidateEnv("EnvVars");

        _databaseConfig = new DatabaseConfig
        (
            host: _configuration["MYSQL_HOST"]!,
            port: _configuration["MYSQL_PORT"]!,
            user: "root",
            password: _configuration["MYSQL_ROOT_PASSWORD"]!,
            database: _configuration["MYSQL_DATABASE"]!
        );
    }
    public void ConfigureServices(IServiceCollection services)
    {
        services.AddControllers();
        services.SetupGithubAuth
        (
            clientId:     _configuration["GITHUB_CLIENT_ID"]!, 
            clientSecret: _configuration["GITHUB_CLIENT_SECRET"]!, 
            redirectUri:  "/api/auth/github/callback"
        );
        services.AddSingleton<DatabaseConfig>(_databaseConfig);
    }

    public void Configure(IApplicationBuilder app, IWebHostEnvironment _)
    {
        app.UseMiddleware<GlobalExceptionHandler>();
        app.UseRouting();
        app.UseAuthentication();
        app.UseAuthorization();
        app.UseEndpoints(endpoints =>
        {
            endpoints.MapControllers();
        });
    }
}
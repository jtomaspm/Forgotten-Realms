using ApiUtils.Auth.GitHub;
using ApiUtils.Configuration;
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

        _databaseConfig = new DatabaseConfig(
            _configuration["MYSQL_HOST"]!,
            _configuration["MYSQL_PORT"]!,
            "root",
            _configuration["MYSQL_ROOT_PASSWORD"]!,
            _configuration["MYSQL_DATABASE"]!);
    }
    public void ConfigureServices(IServiceCollection services)
    {
        services.AddControllers();
        services.SetupGithubAuth(_configuration["GITHUB_CLIENT_ID"]!, _configuration["GITHUB_CLIENT_SECRET"]!, "/api/auth/github/callback");
    }

    public void Configure(IApplicationBuilder app, IWebHostEnvironment _)
    {
        app.UseRouting();
        app.UseAuthentication();
        app.UseAuthorization();
        app.UseEndpoints(endpoints =>
        {
            endpoints.MapControllers();
        });
    }
}
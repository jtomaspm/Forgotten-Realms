using ApiUtils.Configuration;
using ApiUtils.Middleware;
using Dapper;
using Database;
using Database.Application;
using Database.Application.TypeHandlers;

namespace Backend;

public class Startup
{
    private readonly IConfiguration _configuration;
    private readonly DatabaseConfig _databaseConfig;
    private readonly IDatabaseFactory<ApplicationDatabase> _databaseFactory;
    public Startup(IConfiguration configuration)
    {
        _configuration = configuration ?? throw new ArgumentNullException(nameof(configuration));
        _configuration.ValidateEnv("EnvVars");

        _databaseConfig = new DatabaseConfig
        (
            jwtIssuer: _configuration["JWT_ISSUER"]!,
            jwtSecret: _configuration["JWT_SECRET"]!,
            host: _configuration["MYSQL_HOST"]!,
            port: _configuration["MYSQL_PORT"]!,
            user: "root",
            password: _configuration["MYSQL_ROOT_PASSWORD"]!,
            database: _configuration["MYSQL_DATABASE"]!
        );
        _databaseFactory = new DatabaseFactory<ApplicationDatabase>(_databaseConfig);
        SqlMapper.AddTypeHandler(new RoleTypeHandler());
    }
    public void ConfigureServices(IServiceCollection services)
    {
        services.AddControllers();
        services.AddSingleton(_databaseConfig);
        services.AddSingleton(_databaseFactory);
    }

    public void Configure(IApplicationBuilder app, IWebHostEnvironment _)
    {
        app.UseMiddleware<GlobalExceptionHandler>();
        app.UseRouting();
        app.UseEndpoints(endpoints =>
        {
            endpoints.MapControllers();
        });
    }
}
using ApiUtils.Logging;

namespace Backend;

public static class Program
{
    public static void Main(string[] args)
    {
        var builder = CreateHostBuilder(args);
        var host = builder.Build();
        host.Run();
    }

    private static IHostBuilder CreateHostBuilder(string[] args) 
    {
        var hostBuilder = Host.CreateDefaultBuilder(args)
            .SetupDefaultLogging(filename: "backend", useConsole: false)
            .ConfigureAppConfiguration((_, config) => 
            {
                config
                    .SetBasePath(Directory.GetCurrentDirectory())
                    .AddJsonFile("appsettings.json")
                    .AddEnvironmentVariables();  
            })
            .ConfigureWebHostDefaults(webBuilder =>
            {
                webBuilder.UseStartup<Startup>();
                webBuilder.UseUrls();
            });
        return hostBuilder;
    }
        
}
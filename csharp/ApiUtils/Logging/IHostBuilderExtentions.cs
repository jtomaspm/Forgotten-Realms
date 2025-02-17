using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Serilog;

namespace ApiUtils.Logging;

public static class IHostBuilderExtentions
{
    public static IHostBuilder SetupDefaultLogging(this IHostBuilder builder)
    {
        return builder
            .ConfigureLogging(logging => 
            {
                logging.ClearProviders();
                logging.AddSerilog(LoggerFactory.DefaultLogger());
                logging.AddConsole();
            });
    }
}

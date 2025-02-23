using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Serilog;

namespace ApiUtils.Logging;

public static class IHostBuilderExtensions
{
    public static IHostBuilder SetupDefaultLogging(this IHostBuilder builder, string filename, bool useConsole) =>
        builder
            .ConfigureLogging(logging => 
            {
                logging.ClearProviders();
                logging.AddSerilog(LoggerFactory.DefaultLogger(filename));
                if (useConsole) logging.AddConsole();
            });
}

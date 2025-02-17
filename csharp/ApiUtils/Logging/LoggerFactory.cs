using Serilog;
using Serilog.Events;
using Serilog.Formatting.Json;

namespace ApiUtils.Logging;

public class LoggerFactory
{

    public static ILogger DefaultLogger() 
    {
        return new LoggerConfiguration()
            .MinimumLevel.Debug()
            .WriteTo.File(
                new JsonFormatter(),
                "/var/log/log-.log",
                rollingInterval: RollingInterval.Day,
                restrictedToMinimumLevel: LogEventLevel.Information
            )
            .CreateLogger();
    }

}

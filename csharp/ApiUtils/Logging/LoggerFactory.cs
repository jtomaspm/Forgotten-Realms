using Serilog;
using Serilog.Events;
using Serilog.Formatting.Json;

namespace ApiUtils.Logging;

public class LoggerFactory
{

    public static ILogger DefaultLogger(string filename) =>
        new LoggerConfiguration()
            .MinimumLevel.Debug()
            .WriteTo.File(
                new JsonFormatter(),
                $"/var/log/{filename}-.log",
                rollingInterval: RollingInterval.Day,
                restrictedToMinimumLevel: LogEventLevel.Information
            )
            .CreateLogger();

}

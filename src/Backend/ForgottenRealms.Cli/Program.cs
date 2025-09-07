using ForgottenRealms.Simulation;
using ForgottenRealms.Simulation.Models;

namespace ForgottenRealms.Cli;

class Program
{
    static void Main(string[] args)
    {
        var simulationSettings = new SimulationSettings
        {
            TicksPerSecond = 120
        };
        var simulation = new Runtime(simulationSettings);
        using var cts = new CancellationTokenSource();
        Console.CancelKeyPress += (sender, eventArgs) =>
        {
            eventArgs.Cancel = true;
            cts.Cancel();
        };
        simulation.Run(cts.Token);
    }
}

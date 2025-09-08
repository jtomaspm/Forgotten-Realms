using ForgottenRealms.Simulation;
using ForgottenRealms.Simulation.SimulationModels;

namespace ForgottenRealms.Cli;

class Program
{
    static void Main(string[] args)
    {
        var simulationSettings = new SimulationSettings
        {
            TicksPerSecond = 60
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

using ForgottenRealms.Simulation.Models;

namespace ForgottenRealms.Simulation;

public class Runtime(SimulationSettings settings)
{
    private readonly SimulationSettings _settings = settings;
    public void Run(CancellationToken cancellationToken)
    {
        var serverTick = new ServerTick(_settings.TicksPerSecond);
        while (!cancellationToken.IsCancellationRequested)
        {
            
            serverTick.WaitForNextTick(cancellationToken);
        }
    }
}
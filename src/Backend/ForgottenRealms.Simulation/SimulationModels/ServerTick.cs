namespace ForgottenRealms.Simulation.SimulationModels;

public class ServerTick(int ticksPerSecond)
{
    public readonly TimeSpan TickDuration = TimeSpan.FromMilliseconds(1000 / ticksPerSecond);
    public DateTime CurrentTick { get; private set; } = DateTime.UtcNow;
    public void WaitForNextTick(CancellationToken cancellationToken)
    {
        var waitTime = TickDuration - (DateTime.UtcNow - CurrentTick);
        if (waitTime > TimeSpan.Zero)
        {
            Console.WriteLine($"Waiting for next tick: {waitTime.TotalMilliseconds} ms");
            Task.Delay(waitTime, cancellationToken).Wait(cancellationToken);
        }
        CurrentTick = DateTime.UtcNow;
    }
}
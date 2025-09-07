namespace ForgottenRealms.Simulation.Models;

public class ServerTick(int ticksPerSecond)
{
    public readonly TimeSpan TickDuration = TimeSpan.FromMilliseconds(1000 / ticksPerSecond);
    public DateTime LastTick { get; private set; } = DateTime.UtcNow;
    public void WaitForNextTick(CancellationToken cancellationToken)
    {
        var waitTime = TickDuration - (DateTime.UtcNow - LastTick);
        if (waitTime > TimeSpan.Zero)
        {
            Task.Delay(waitTime, cancellationToken).Wait(cancellationToken);
        }
        LastTick = DateTime.UtcNow;
    }
}
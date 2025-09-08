using ForgottenRealms.Simulation.SimulationModels;

namespace ForgottenRealms.Simulation.GameLogic.Interfaces;

public interface IGameEntity
{
    public Task Update(ServerTick serverTick, CancellationToken cancellationToken);
}

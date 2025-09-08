using ForgottenRealms.Simulation.GameLogic.Factories;
using ForgottenRealms.Simulation.GameLogic.Interfaces;
using ForgottenRealms.Simulation.SimulationModels;

namespace ForgottenRealms.Simulation;

public class Runtime(SimulationSettings settings)
{
    private readonly SimulationSettings _settings = settings;
    private readonly List<IGameEntity> _gameEntities = [];
    public void Run(CancellationToken cancellationToken)
    {
        var serverTick = new ServerTick(_settings.TicksPerSecond);
        _gameEntities.AddRange(VillageFactory.RandomVillages(1000*1000));

        while (!cancellationToken.IsCancellationRequested)
        {
            Parallel.ForEach(_gameEntities, entity => entity.Update(serverTick, cancellationToken));
            
            serverTick.WaitForNextTick(cancellationToken);
        }
    }
}
using ForgottenRealms.Simulation.GameLogic.Interfaces;
using ForgottenRealms.Simulation.GameLogic.Models;
using ForgottenRealms.Simulation.SimulationModels;

namespace ForgottenRealms.Simulation.GameLogic.GameEntities;

public class Village : IGameEntity
{
    public Coordinates Coordinates { get; set; } = new() { X = 0, Y = 0 };
    public Resources Resources { get; set; } = new(){ Wood = 750, Clay = 750, Iron = 750 };
    public Resources ResourcesPerHour { get; set; } = new() { Wood = 20, Clay = 20, Iron = 20 };
    public Dictionary<BuildingEnum, int> BuildingLevels { get; set; } = new()
    {
        { BuildingEnum.Headquarters, 1 },
        { BuildingEnum.Woodcutter, 0 },
        { BuildingEnum.ClayPit, 0 },
        { BuildingEnum.IronMine, 0 },
    };
    public Task Update(ServerTick tick, CancellationToken cancellationToken)
    {
        var interval = tick.TickDuration.TotalHours;
        Resources.Wood += ResourcesPerHour.Wood * interval;
        Resources.Clay += ResourcesPerHour.Clay * interval;
        Resources.Iron += ResourcesPerHour.Iron * interval;
        return Task.CompletedTask;
    }

    public override string ToString() => $"Village [{Coordinates.X}|{Coordinates.Y}] - Wood: {Resources.Wood}, Clay: {Resources.Clay}, Iron: {Resources.Iron}";
}

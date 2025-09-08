using ForgottenRealms.Simulation.GameLogic.GameEntities;

namespace ForgottenRealms.Simulation.GameLogic.Factories;

public class VillageFactory
{
    public static List<Village> RandomVillages(int amount = 1)
    {
        var villages = new List<Village>();
        var rand = new Random();
        for (var i = 0; i < amount; i++)
        {
            villages.Add(new Village
            {
                Coordinates = new()
                {
                    X = rand.Next(-500, 500),
                    Y = rand.Next(-500, 500)
                },
                ResourcesPerHour = new()
                {
                    Wood = rand.Next(100, 400),
                    Clay = rand.Next(1000, 2000),
                    Iron = rand.Next(5000, 10000),
                },
            });
        }
        return villages;
    }
}

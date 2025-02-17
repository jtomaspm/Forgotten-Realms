namespace Database.ApplicationDatabase.Models;

public class World
{
    public Guid Id;
    public required string Name;
    public required string Database;
    public float Speed;
    public required string GameVersion;
    public DateTime CreatedAt;
}




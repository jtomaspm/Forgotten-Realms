namespace Database;

public class DatabaseFactory<TDatabase>(DatabaseConfig config) : IDatabaseFactory<TDatabase> where TDatabase : Database, new()
{
    public async Task<TDatabase> CreateDatabase()
        => await Database.CreateAsync<TDatabase>(config);
}

public interface IDatabaseFactory<TDatabase> where TDatabase : Database
{
    Task<TDatabase> CreateDatabase();
}
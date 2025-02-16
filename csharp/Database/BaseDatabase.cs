using MySql.Data.MySqlClient;

namespace Database;

public abstract class BaseDatabase : IDisposable
{
    protected readonly DatabaseConfig _config;
    protected readonly MySqlConnection _connection;

    public BaseDatabase(DatabaseConfig config)
    {
        _config = config;
        _connection = new MySqlConnection(_config.ConnectionString);
    }

    public void Dispose()
    {
        _connection.Dispose();
    }
}

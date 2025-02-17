using MySql.Data.MySqlClient;

namespace Database;

public class Database : IDisposable
{
    protected readonly DatabaseConfig _config;
    protected readonly MySqlConnection _connection;

    public Database(DatabaseConfig config)
    {
        _config = config;
        _connection = new MySqlConnection(_config.ConnectionString);
    }

    public MySqlConnection GetConnection() 
    {
        return _connection;
    }

    public void Dispose()
    {
        _connection.Dispose();
    }
}

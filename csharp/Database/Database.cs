using System.Dynamic;
using MySql.Data.MySqlClient;

namespace Database;

public class Database : IDisposable
{
    protected DatabaseConfig? _config;
    protected MySqlConnection? _connection;

    private Database(){}

    public Database(DatabaseConfig config)
    {
        _config = config;
        Database.SetupDatabase(this);
    }

    private static async Task SetupDatabaseAsync(Database database) 
    {
        if (database._config is null) throw new ArgumentNullException("Attempting to Setup a database that is not configured.");

        database._connection = new MySqlConnection(database._config.ConnectionString);
        await database._connection.OpenAsync();
    }

    private static void SetupDatabase(Database database) 
    {
        if (database._config is null) throw new ArgumentNullException("Attempting to Setup a database that is not configured.");

        database._connection = new MySqlConnection(database._config.ConnectionString);
        database._connection.Open();
    }

    public static async Task<Database> CreateAsync(DatabaseConfig config)
    {
        var database = new Database
        {
            _config = config
        };

        await Database.SetupDatabaseAsync(database);

        return database;       
    }

    public MySqlConnection GetConnection() 
    {
        if (_connection is null) 
        {
            Database.SetupDatabase(this);
        }
        return _connection!;
    }

    public async Task<MySqlConnection> GetConnectionAsync() 
    {
        if (_connection is null) 
        {
            await Database.SetupDatabaseAsync(this);
        }
        return _connection!;
    }

    public DatabaseConfig GetConfig() => this._config ?? throw new ArgumentNullException("This database should be configured.");

    public void Dispose()
    {
        _connection?.Dispose();
        GC.SuppressFinalize(this);
    }
}

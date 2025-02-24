using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Security.Cryptography;
using System.Text;
using Microsoft.IdentityModel.Tokens;
using MySql.Data.MySqlClient;

namespace Database;

public abstract class Database : IDisposable
{
    protected DatabaseConfig? _config;
    protected MySqlConnection? _connection;
    protected MySqlTransaction? _transaction;
    internal bool InTrasaction { get => _transaction is not null; }
    internal Database(){}

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

    public static async Task<TDatabase> CreateAsync<TDatabase>(DatabaseConfig config) where TDatabase : Database, new()
    {
        var database = new TDatabase
        {
            _config = config
        };

        await Database.SetupDatabaseAsync(database);

        return database;       
    }

    internal MySqlConnection GetConnection() 
    {
        if (_connection is null) 
        {
            Database.SetupDatabase(this);
        }
        return _connection!;
    }

    internal async Task<MySqlConnection> GetConnectionAsync() 
    {
        if (_connection is null) 
        {
            await Database.SetupDatabaseAsync(this);
        }
        return _connection!;
    }

    internal async Task ExecuteInTransaction(Func<MySqlConnection,MySqlTransaction,Task> callback)
    {
        if (_connection is null) await GetConnectionAsync();
        _transaction ??= _connection!.BeginTransaction() ?? throw new Exception("Error creating database transaction");
        try 
        {
            await callback(_connection!, _transaction);
            await _transaction.CommitAsync();
            await _transaction.DisposeAsync();
            _transaction = null;
        } 
        catch (Exception e) 
        {
            await _transaction!.RollbackAsync();
            await _transaction.DisposeAsync();
            _transaction = null;
            throw new Exception("Create Account Transaction failed", e);
        }
    }


    public async Task ExecuteInTransaction(Func<MySqlTransaction,Task> callback)
    {
        await ExecuteInTransaction(async (connection, transaction) => {
            await callback(transaction);
        });
    }

    internal async Task<TResult> ExecuteInTransaction<TResult>(Func<MySqlConnection,MySqlTransaction,Task<TResult>> callback)
    {
        if (_connection is null) await GetConnectionAsync();
        _transaction ??= _connection!.BeginTransaction() ?? throw new Exception("Error creating database transaction");
        try 
        {
            var result = await callback(_connection!, _transaction);
            await _transaction.CommitAsync();
            await _transaction.DisposeAsync();
            _transaction = null;
            return result;
        } 
        catch (Exception e) 
        {
            await _transaction!.RollbackAsync();
            await _transaction.DisposeAsync();
            _transaction = null;
            throw new Exception("Create Account Transaction failed", e);
        }
    }

    public async Task<TResult> ExecuteInTransaction<TResult>(Func<MySqlTransaction,Task<TResult>> callback) =>
        await ExecuteInTransaction(async (connection, transaction) => await callback(transaction));

    public DatabaseConfig GetConfig() => _config ?? throw new ArgumentNullException("This database should be configured.");

    internal string GenerateJwtToken(Guid sessionId, Guid accountId)
    {
        var securityKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(GetConfig().JwtSecret));
        var credentials = new SigningCredentials(securityKey, SecurityAlgorithms.HmacSha256);
        
        var claims = new[]
        {
            new Claim("sessionId", sessionId.ToString()),
            new Claim("accountId", accountId.ToString())
        };
        
        var token = new JwtSecurityToken(
            issuer: GetConfig().JwtIssuer,
            claims: claims,
            expires: DateTime.UtcNow.AddDays(7),
            signingCredentials: credentials
        );

        return new JwtSecurityTokenHandler().WriteToken(token);
    }

    public static string SHA256Token(string jwtToken) =>
        Convert.ToBase64String(SHA256.HashData(Encoding.UTF8.GetBytes(jwtToken)));

    public void Dispose()
    {
        _connection?.Dispose();
        GC.SuppressFinalize(this);
    }
}

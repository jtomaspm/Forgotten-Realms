namespace Database;

public class DatabaseConfig(string host, string port, string user, string password, string database, string jwtIssuer, string jwtSecret)
{
    public readonly string Host = host;
    public readonly string Port = port;
    public readonly string User = user;
    public readonly string Password = password;
    public readonly string Database = database;
    public readonly string JwtIssuer = jwtIssuer;
    public readonly string JwtSecret = jwtSecret;

    public string ConnectionString
    {
        get
        {
            return $"Server={Host};Port={Port};Database={Database};User={User};Password={Password};";
        }
    }

    public DatabaseConfig WithDatabase(string newDatabase) => new DatabaseConfig
    (
        host: Host,
        port: Port,
        user: User,
        password: Password,
        database: newDatabase,
        jwtIssuer: JwtIssuer,
        jwtSecret: JwtSecret
    );
}
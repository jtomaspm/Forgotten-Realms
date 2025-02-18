namespace Database;

public class DatabaseConfig
{
    public readonly string Host;
    public readonly string Port;
    public readonly string User;
    public readonly string Password;
    public readonly string Database;

    public DatabaseConfig(string host, string port, string user, string password, string database)
    {
        Host = host;
        Port = port;
        User = user;
        Password = password;
        Database = database;
    }

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
        database: newDatabase
    );
}
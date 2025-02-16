namespace Database;

public class DatabaseConfig
{
    public readonly string Host;
    public readonly string Port;
    public readonly string User;
    public readonly string Password;
    public readonly string Database;
    public readonly string ConnectionString;
    
    public DatabaseConfig(string host, string port, string user, string password, string database)
    {
        Host = host;
        Port = port;
        User = user;
        Password = password;
        Database = database;
        ConnectionString = $"Server={Host};Port={Port};Database={Database};User={User};Password={Password};";
    }
}
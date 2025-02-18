namespace Database;

public class DatabaseConfig
{
    public required string Host;
    public required string Port;
    public required string User;
    public required string Password;
    public required string Database;
    public string ConnectionString
    {
        get
        {
            return $"Server={Host};Port={Port};Database={Database};User={User};Password={Password};";
        }
    }
}
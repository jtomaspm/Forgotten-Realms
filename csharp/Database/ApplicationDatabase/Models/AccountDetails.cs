namespace Database.ApplicationDatabase.Models;

public class AccountDetails: Account
{
    public required AccountProperties AccountProperties;
    public Login? LastLogin;
    public Session? Session;
    public required IEnumerable<World> Worlds;
}

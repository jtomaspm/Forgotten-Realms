using Database.ApplicationDatabase.Models;

namespace Database.ApplicationDatabase.Services;

public class AccountService
{
    public static Account CreateAccout(Database database, string source, string externalId, string name, string email, Role role)
    {
        var account = new Account
        {
            Name = name,
            Email = email,
            Source = source,
            ExternalId = externalId,
            Role = role
        };

        var connetion = database.GetConnection();


        return account;
    }
}

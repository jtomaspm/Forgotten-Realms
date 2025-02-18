using Database.ApplicationDatabase.Models;

namespace Database.ApplicationDatabase.Services;

public class AccountService
{
    public static async Task<Account> CreateAccount(Database database, string source, string externalId, string name, string email, Role role)
    {
        var account = new Account
        {
            Name = name,
            Email = email,
            Source = source,
            ExternalId = externalId,
            Role = role
        };

        var connection = await database.GetConnectionAsync();
        using (var transaction = connection.BeginTransaction() ?? throw new Exception("Error creating database transaction"))
        {
            try 
            {
                var cmd = connection.CreateCommand();
                cmd.CommandText = "SELECT `Id`, `Name` FROM `Accounts` WHERE Email=@email";
                cmd.Parameters.AddWithValue("email", "joaotomachado@proton.me");

                using (var reader = await cmd.ExecuteReaderAsync())
                while (await reader.ReadAsync())
                {
                        var accountId = reader.GetGuid(0);
                        var accountName = reader.GetString(1);
                        Console.WriteLine($"$$$$$$$$$$$$[Account]$$$$$$$$$\nId:{accountId}\nName:{accountName}");
                }

                await transaction.CommitAsync();
            } 
            catch (Exception e) 
            {
                await transaction.RollbackAsync();
                throw new Exception("Create Account Transaction failed", e);
            }
        }

        return account;
    }
}

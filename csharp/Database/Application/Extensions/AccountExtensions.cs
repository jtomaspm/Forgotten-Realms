using Database.Application.Models;
using Dapper;


namespace Database.Application.Extensions;

public static class AccountExtensions
{
    public static async Task<Account?> GetAccountById(this ApplicationDatabase database, Guid id) =>
        await (await database.GetConnectionAsync())
            .QueryFirstOrDefaultAsync<Account>(@$"
                SELECT * 
                FROM `Accounts`
                WHERE `Id`=@id
                LIMIT 1
            ", new {id});

    public static async Task<Account?> GetAccountByExternalId(this ApplicationDatabase database, string externalId, string source) =>
        await (await database.GetConnectionAsync())
            .QueryFirstOrDefaultAsync<Account>(@$"
                SELECT * 
                FROM `Accounts`
                WHERE `ExternalId`=@externalId AND `Source`=@source
                LIMIT 1
            ", new {externalId, source});

    public static async Task<AccountDetails> CreateAccount(this ApplicationDatabase database, string source, string externalId, string? name, string email, Role role)
    {
        if (!database.InTrasaction) 
            return await database.ExecuteInTransaction(
                async (_, _) => 
                    await database.CreateAccount(source, externalId, name, email, role));
        
        var connection = await database.GetConnectionAsync();
        
        var id = Guid.NewGuid();
        var date = DateTime.UtcNow;

        var rowsAffected = await connection.ExecuteAsync(@"
            INSERT INTO `Accounts`
                (`Id`, `Name`, `Email`, `Role`, `ExternalId`, `Source`, `CreatedAt`, `UpdatedAt`)
            VALUES
                (@id, @name, @email, @role, @externalId, @source, @createdAt, @updatedAt)
        ", new {id, name, email, role, externalId, source, createdAt=date, updatedAt=date});

        if (rowsAffected != 1) throw new Exception("Error inserting new account in database.");

        var accountProperties = await database.CreateAccountProperties(id, true, false);

        return new() 
        {
            Id = id,
            Name = name,
            Email = email,
            Role = role,
            ExternalId = externalId,
            Source = source,
            CreatedAt = date,
            UpdatedAt = date,
            AccountProperties = accountProperties,
            LastLogin = null,
            Session = null,
            Worlds = []
        };
    }
}

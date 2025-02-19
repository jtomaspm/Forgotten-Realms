using Database.ApplicationDatabase.Models;
using Database.CommandBuilder.CommandBuilder;
using MySql.Data.MySqlClient;

namespace Database.ApplicationDatabase.Services;

public static class AccountService
{
    public static async Task<Account?> GetById(Guid id, Database database)
    {
        var connection = await database.GetConnectionAsync();
        var cmd = connection.CreateCommand();
        
        cmd.CommandText = @"
            SELECT 
                `Id`, 
                `ExternalId`, 
                `Source`, 
                `Name`, 
                `Email`, 
                `Role`, 
                `CreatedAt`, 
                `UpdatedAt` 
            FROM `Accounts` 
            WHERE `Id`=@id";

        cmd.Parameters.AddWithValue("id", id);

        using var reader = await cmd.ExecuteReaderAsync();
        if (await reader.ReadAsync())
        {
            return new Account() 
            {
                Id = reader.GetGuid(0),
                ExternalId = reader.GetString(1),
                Source = reader.GetString(2),
                Name = reader.GetString(3),
                Email = reader.GetString(4),
                Role = Role.FromName(reader.GetString(4)),
                CreatedAt = reader.GetDateTime(5),
                UpdatedAt = reader.GetDateTime(6),
            };
        }

        return null;
    }

    public static async Task<Account?> GetByExternalId(string externalId, string source, Database database)
    {
        var connection = await database.GetConnectionAsync();
        var cmd = CommandBuilder.Select()
        var cmd = connection.CreateCommand();
        
        cmd.CommandText = @"
            SELECT 
                `Id`, `ExternalId`, `Source`, `Name`, `Email`, `Role`, `CreatedAt`, `UpdatedAt` 
            FROM `Accounts` 
            WHERE `ExternalId`=@externalId
              AND `Source`=@source;";

        cmd.Parameters.AddWithValue("externalId", externalId);
        cmd.Parameters.AddWithValue("source", source);

        using var reader = await cmd.ExecuteReaderAsync();
        if (await reader.ReadAsync())
        {
            return new Account() 
            {
                Id = reader.GetGuid(0),
                ExternalId = reader.GetString(1),
                Source = reader.GetString(2),
                Name = reader.GetString(3),
                Email = reader.GetString(4),
                Role = Role.FromName(reader.GetString(4)),
                CreatedAt = reader.GetDateTime(5),
                UpdatedAt = reader.GetDateTime(6),
            };
        }

        return null;
    }

    private static async Task<int> InsertAccount(MySqlConnection connection, MySqlTransaction _, string source, string externalId, string name, string email, Role role)
    {
        var cmd = connection.CreateCommand();
        cmd.CommandText = @"
            INSERT INTO `Accounts` (
                `Name`, `Email`, `Role`, `ExternalId`, `Source`
            ) VALUES (
                @name, @email, @role, @externalId, @source
            );";

        cmd.Parameters.AddWithValue("name", name);
        cmd.Parameters.AddWithValue("email", email);
        cmd.Parameters.AddWithValue("role", role.Name);
        cmd.Parameters.AddWithValue("externalId", externalId);
        cmd.Parameters.AddWithValue("source", source);

        return await cmd.ExecuteNonQueryAsync();
    }

    public static async Task<Account> CreateAccount(string source, string externalId, string name, string email, Role role, Database database)
    {
        var insertedRows = await database.ExecuteInTransaction<int>(
            (connection, transaction) => 
                InsertAccount(connection, transaction, source, externalId, name, email, role));
        if (insertedRows == 1) {
            var account = await GetByExternalId(externalId, source, database);
            if (account is not null) return account;
        }
        throw new Exception("Uncaught error while creating account.");
    }
}

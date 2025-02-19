using Database.ApplicationDatabase.Models;
using Database.CommandBuilder;
using MySql.Data.MySqlClient;

namespace Database.ApplicationDatabase.Extensions;

public static class AccountExtensions
{
    private static async IAsyncEnumerable<Account> GetWithParams(IEnumerable<string> conditions, Dictionary<string, object> parameters, Database database) 
    {
        var connection = await database.GetConnectionAsync();
        var cmd = connection.CreateCommand().Select(
            ["Id", "ExternalId, Source, Name, Email, Role, CreatedAt, UpdatedAt"], 
            "Accounts", 
            conditions, 
            parameters);
        
        using var reader = await cmd.ExecuteReaderAsync();
        while (await reader.ReadAsync())
        {
            yield return new Account() 
            {
                Id = reader.GetGuid(0),
                ExternalId = reader.GetString(1),
                Source = reader.GetString(2),
                Name = reader.GetString(3),
                Email = reader.GetString(4),
                Role = Role.FromName(reader.GetString(5)),
                CreatedAt = reader.GetDateTime(6),
                UpdatedAt = reader.GetDateTime(7),
            };
        }
    }

    public static async Task<Account?> GetAccountById(this ApplicationDatabase database, Guid id)
    {
        var query_result = GetWithParams(["Id=@id"], new(){{ "id", id }}, database);
        await foreach (var account in query_result)
        {
            return account;
        }
        return null;
    }

    public static async Task<Account?> GetAccountByExternalId(this ApplicationDatabase database, string externalId, string source)
    {
        var query_result = GetWithParams(
            ["ExternalId`=@externalId", "`Source`=@source"], 
            new(){{ "externalId", externalId }, { "source", source }}, 
            database);
        await foreach (var account in query_result)
        {
            return account;
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

    public static async Task<Account> CreateAccount(this ApplicationDatabase database, string source, string externalId, string name, string email, Role role)
    {
        var insertedRows = await database.ExecuteInTransaction<int>(
            (connection, transaction) => 
                InsertAccount(connection, transaction, source, externalId, name, email, role));
        if (insertedRows == 1) {
            var account = await database.GetAccountByExternalId(externalId, source);
            if (account is not null) return account;
        }
        throw new Exception("Uncaught error while creating account.");
    }
}

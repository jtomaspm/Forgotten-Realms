using Database.ApplicationDatabase.Models;
using Database.CommandBuilder;
using MySql.Data.MySqlClient;


namespace Database.ApplicationDatabase.Extensions;

public static class AccountExtensions
{
    public static async Task<Account?> GetAccountById(this ApplicationDatabase database, Guid id)
        => await GetWithParams
        (
            ["Id=@id"], 
            new()
            {
                { "id", id }
            }, 
            database
        )
        .FirstOrDefaultAsync();

    public static async Task<Account?> GetAccountByExternalId(this ApplicationDatabase database, string externalId, string source)
        => await GetWithParams
        (
            ["ExternalId`=@externalId", "`Source`=@source"], 
            new()
            {
                { "externalId", externalId }, { "source", source }
            }, 
            database
        )
        .FirstOrDefaultAsync();
    public static async Task<Account> CreateAccount(this ApplicationDatabase database, string source, string externalId, string name, string email, Role role)
        => await database.ExecuteInTransaction<Account>
        (
            async (connection, transaction) => 
            {
                Account account;
                if (1 == await InsertAccount(connection, null, source, externalId, name, email, role))
                    account = await database.GetAccountByExternalId(externalId, source) ?? throw new Exception("Uncaught error while creating account.");
                else
                    throw new Exception("Error inserting new account in database.");

                return account;
            }
        );
    private static async IAsyncEnumerable<Account> GetWithParams(IEnumerable<string> conditions, Dictionary<string, object> parameters, Database database) 
    {
        var connection = await database.GetConnectionAsync();
        var cmd = connection.CreateCommand().Select
        (
            fields: ["Id", "ExternalId, Source, Name, Email, Role, CreatedAt, UpdatedAt"], 
            table: "Accounts", 
            conditions: conditions, 
            parameters: parameters
        );
        
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
    private static async Task<int> InsertAccount(MySqlConnection connection, MySqlTransaction? _, string source, string externalId, string name, string email, Role role)
        => await connection
            .CreateCommand()
            .Insert
            (
                table: "Accounts",
                fields: ["Name", "Email", "Role", "ExternalId", "Source"],
                values: [["@name", "@email", "@role", "@externalId", "@source"]], 
                parameters: new()
                {
                    {"name", name},
                    {"email", email},
                    {"role", role.Name},
                    {"externalId", externalId},
                    {"source", source},
                }
            )
            .ExecuteNonQueryAsync();
}

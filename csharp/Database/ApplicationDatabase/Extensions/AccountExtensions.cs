using System.Reflection.Metadata.Ecma335;
using Database.ApplicationDatabase.Models;
using Org.BouncyCastle.Asn1.Cms;


namespace Database.ApplicationDatabase.Extensions;

public static class AccountExtensions
{
    public static async Task<Account?> GetAccountById(this ApplicationDatabase database, Guid id)
        => await GetWithParams
        (
            conditions: ["Id=@id"], 
            parameters: new() {{ "id", id }}, 
            database:   database
        )
        .FirstOrDefaultAsync();
    public static async Task<Account?> GetAccountByExternalId(this ApplicationDatabase database, string externalId, string source)
        => await GetWithParams
        (
            conditions: ["ExternalId`=@externalId", "`Source`=@source"], 
            parameters: new() {{ "externalId", externalId }, { "source", source }}, 
            database: database
        )
        .FirstOrDefaultAsync();
    public static async Task<AccountDetails> CreateAccount(this ApplicationDatabase database, string source, string externalId, string name, string email, Role role)
        => await database.ExecuteInTransaction<AccountDetails>
        (
            async (connection, transaction) => 
            {
                var cmd = (await database.Insert())
                    .Table("Accounts")
                    .AddFields(["Name", "Email", "Role", "ExternalId", "Source"])
                    .AddValues([["@name", "@email", "@role", "@externalId", "@source"]])
                    .SetParameters(new()
                        {
                            {"name", name},
                            {"email", email},
                            {"role", role.Name},
                            {"externalId", externalId},
                            {"source", source},
                        })
                    .Build();

                Account account;
                var insertResult = await cmd.ExecuteNonQueryAsync();
                if (insertResult == 1)
                    account = await database.GetAccountByExternalId(externalId, source) 
                        ?? throw new Exception("Uncaught error while creating account.");
                else
                    throw new Exception("Error inserting new account in database.");

                var accountProperties = await database.CreateAccountProperties(account.Id, false, false);

                var accountDetails = (AccountDetails) account;
                accountDetails.AccountProperties = accountProperties;
                accountDetails.Worlds = [];

                return accountDetails;
            }
        );
    private static async IAsyncEnumerable<Account> GetWithParams(IEnumerable<string> conditions, Dictionary<string, object> parameters, Database database) 
    {
        var cmd = (await database.Select())
            .AddFields(["Id", "ExternalId, Source, Name, Email, Role, CreatedAt, UpdatedAt"])
            .Table("Accounts")
            .AddConditions(conditions)
            .SetParameters(parameters)
            .Build();
        
        using var reader = await cmd.ExecuteReaderAsync();
        while (await reader.ReadAsync())
        {
            yield return new Account() 
            {
                Id = reader.GetGuid(0),
                ExternalId = reader.IsDBNull(1) ? null : reader.GetString(1),
                Source = reader.IsDBNull(2) ? null : reader.GetString(2),
                Name = reader.GetString(3),
                Email = reader.GetString(4),
                Role = Role.FromName(reader.GetString(5)),
                CreatedAt = reader.GetDateTime(6),
                UpdatedAt = reader.GetDateTime(7),
            };
        }
    }
}

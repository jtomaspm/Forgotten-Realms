using Database.Application.Models;

namespace Database.Application.Extensions;

public static class AccountPropertiesExtensions
{
    public static async Task<AccountProperties?> GetAccountPropertiesById(this ApplicationDatabase database, Guid accountId)
        => await GetWithParams
        (
            conditions: ["AccountId=@accountId"], 
            parameters: new() { { "accountId", accountId } },
            database:   database
        )
        .FirstOrDefaultAsync();

    public static async Task<AccountProperties> CreateAccountProperties(this ApplicationDatabase database, Guid accountId, bool emailVerified, bool sendEmailNotifications)
    {
        var cmd = (await database.Insert())
            .Table("AccountProperties")
            .AddFields(["AccountId", "VerificationToken", "TokenExpiresAt", "EmailVerified", "SendEmailNotifications"])
            .AddValues([["@accountId", "@verificationToken", "@tokenExpiresAt", "@emailVerified", "@sendEmailNotifications"]])
            .SetParameters(new()
                {
                    { "accountId", accountId },
                    { "emailVerified", emailVerified },
                    { "sendEmailNotifications", sendEmailNotifications },
                })
            .Build();

        AccountProperties accountProperties;
        var insertResult = await cmd.ExecuteNonQueryAsync();
        if (insertResult == 1)
            accountProperties = await database.GetAccountPropertiesById(accountId)
                ?? throw new Exception("Uncaught error while creating account properties.");
        else
            throw new Exception("Error inserting new account properties in database.");
        
        return accountProperties;
    }

    private static async IAsyncEnumerable<AccountProperties> GetWithParams(IEnumerable<string> conditions, Dictionary<string, object> parameters, Database database)
    {
        var cmd = (await database.Select())
            .AddFields(["AccountId", "VerificationToken", "TokenExpiresAt", "EmailVerified", "SendEmailNotifications", "CreatedAt", "UpdatedAt"])
            .Table("AccountProperties")
            .AddConditions(conditions)
            .SetParameters(parameters)
            .Build();
        
        using var reader = await cmd.ExecuteReaderAsync();
        while (await reader.ReadAsync())
        {
            yield return new AccountProperties()
            {
                AccountId = reader.GetGuid(0),
                VerificationToken = reader.GetGuid(1),
                TokenExpiresAt = reader.GetDateTime(2),
                EmailVerified = reader.GetBoolean(3),
                SendEmailNotifications = reader.GetBoolean(4),
                CreatedAt = reader.GetDateTime(5),
                UpdatedAt = reader.GetDateTime(6)
            };
        }
    }
}
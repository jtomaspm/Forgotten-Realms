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
        var verificationToken = Guid.NewGuid();
        var date = DateTime.UtcNow;
        var expiresAt = date.AddDays(7);

        var cmd = (await database.Insert())
            .Table("AccountProperties")
            .AddFields(["AccountId", "VerificationToken", "TokenExpiresAt", "EmailVerified", "SendEmailNotifications", "CreatedAt", "UpdatedAt"])
            .AddValues([["@accountId", "@verificationToken", "@expiresAt", "@emailVerified", "@sendEmailNotifications", "createdAt", "updatedAt"]])
            .SetParameters(new()
                {
                    { "accountId", accountId },
                    { "verificationToken", verificationToken },
                    { "expiresAt", expiresAt },
                    { "emailVerified", emailVerified },
                    { "sendEmailNotifications", sendEmailNotifications },
                    { "createdAt", date },
                    { "updatedAt", date },
                })
            .Build();

        var insertResult = await cmd.ExecuteNonQueryAsync();
        if (insertResult != 1)
            throw new Exception("Error inserting new account properties in database.");

        return new() 
        {
            AccountId = accountId,
            VerificationToken = verificationToken,
            TokenExpiresAt = expiresAt,
            EmailVerified = emailVerified,
            SendEmailNotifications = sendEmailNotifications,
            CreatedAt = date,
            UpdatedAt = date,
        };
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
            yield return new()
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
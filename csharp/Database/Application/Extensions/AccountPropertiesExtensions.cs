using Dapper;
using Database.Application.Models;

namespace Database.Application.Extensions;

public static class AccountPropertiesExtensions
{
    public static async Task<AccountProperties?> GetAccountPropertiesById(this ApplicationDatabase database, Guid accountId) => 
        await (await database.GetConnectionAsync())
            .QueryFirstOrDefaultAsync<AccountProperties>(@$"
                SELECT * 
                FROM `AccountProperties`
                WHERE `AccountId`=@accountId
                LIMIT 1
            ", new {accountId});

    public static async Task<AccountProperties> CreateAccountProperties(this ApplicationDatabase database, Guid accountId, bool emailVerified, bool sendEmailNotifications)
    {
        var connection = await database.GetConnectionAsync();
        var verificationToken = Guid.NewGuid();
        var date = DateTime.UtcNow;
        var tokenExpiresAt = date.AddDays(7);

        var insertResult = await connection.ExecuteAsync(@"
            INSERT INTO `AccountProperties`
                (`AccountId`, `VerificationToken`, `TokenExpiresAt`, `EmailVerified`, `SendEmailNotifications`, `CreatedAt`, `UpdatedAt`)
            VALUES
                (@accountId, @verificationToken, @tokenExpiresAt, @emailVerified, @sendEmailNotifications, @createdAt, @updatedAt)
        ", new {accountId, verificationToken, tokenExpiresAt, emailVerified, sendEmailNotifications, createdAt=date, updatedAt=date});

        if (insertResult != 1)
            throw new Exception("Error inserting new account properties in database.");

        return new() 
        {
            AccountId = accountId,
            VerificationToken = verificationToken,
            TokenExpiresAt = tokenExpiresAt,
            EmailVerified = emailVerified,
            SendEmailNotifications = sendEmailNotifications,
            CreatedAt = date,
            UpdatedAt = date,
        };
    }
}
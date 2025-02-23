using Database.Application.Models;

namespace Database.Application.Extensions;

public static class AuthExtensions
{
    public static async Task<Session?> GetSessionById(this ApplicationDatabase database, Guid id)
        => await GetSessionsWithParams(["Id=@id"], new() { { "id", id } }, database).FirstOrDefaultAsync();
    public static async Task<Session?> GetSessionByAccountId(this ApplicationDatabase database, Guid accountId)
        => await GetSessionsWithParams(["AccountId=@accountId"], new() { { "accountId", accountId } }, database).FirstOrDefaultAsync();
    public static async Task<Session?> GetSessionByToken(this ApplicationDatabase database, string token)
        => await GetSessionsWithParams(["Token=@token"], new() { { "token", token } }, database).FirstOrDefaultAsync();
    public static async Task<Session> CreateSession(this ApplicationDatabase database, Guid accountId)
    {
        var session = await database.GetSessionByAccountId(accountId);
        if (session is not null && (await database.DeleteSession(accountId) == 0)) 
            throw new Exception($"Error deleting current account session for AccountId: {accountId}");    

        var sid = Guid.NewGuid();
        var token = database.GenerateJwtToken(sid, accountId);
        var createdAt = DateTime.UtcNow;
        var expiresAt = createdAt + TimeSpan.FromDays(7);

        var cmd = (await database.Insert())
            .Table("Sessions")
            .AddFields(["Id", "AccountId", "Token", "ExpiresAt"])
            .AddValues([["@sessionId", "@accountId", "@token", "@expiresAt", "@createdAt"]])
            .SetParameters(new()
                {
                    { "sessionId", sid },
                    { "accountId", accountId },
                    { "token", token },
                    { "expiresAt", expiresAt },
                    { "createdAt", createdAt }
                })
            .Build();

        var insertResult = await cmd.ExecuteNonQueryAsync();
        if (insertResult != 1)
            throw new Exception("Error inserting new session in database.");

        return new() 
        {
            Id = sid,
            AccountId = accountId,
            Token = token,
            CreatedAt = createdAt,
            ExpiresAt = expiresAt
        };        
    }
    public static async Task<int> DeleteSession(this ApplicationDatabase database, Guid accountId) =>
        await (await database.Delete())
            .Table("Sessions")
            .AddCondition("AccountId=@accountId")
            .SetParameters(new() {{ "accountId", accountId }})
            .Build()
            .ExecuteNonQueryAsync();

    public static async Task<Login?> GetLoginById(this ApplicationDatabase database, Guid loginId)
        => await GetLoginsWithParams(["Id=@loginId"], new() { { "loginId", loginId } }, database).FirstOrDefaultAsync();

    public static async Task<Login?> GetLoginByIpAddress(this ApplicationDatabase database, string ipAddress)
        => await GetLoginsWithParams(["IpAddress=@ipAddress"], new() { { "ipAddress", ipAddress } }, database).FirstOrDefaultAsync();

    public static async Task<Login> CreateLogin(this ApplicationDatabase database, Guid accountId, string ipAddress)
    {
        var id = Guid.NewGuid();
        var createdAt = DateTime.UtcNow;
        var cmd = (await database.Insert())
            .Table("Logins")
            .AddFields(["AccountId", "IpAddress", "CreatedAt"])
            .AddValues([["@accountId", "@ipAddress", "@createdAt"]])
            .SetParameters(new()
                {
                    { "id", id },
                    { "accountId", accountId },
                    { "ipAddress", ipAddress },
                    { "createdAt", createdAt },
                })
            .Build();

        var insertResult = await cmd.ExecuteNonQueryAsync();
        if (insertResult != 1)
            throw new Exception("Error inserting new login record in database.");

        return new() { Id=id, AccountId = accountId, CreatedAt = createdAt, IpAddress = ipAddress };
    }

    private static async IAsyncEnumerable<Session> GetSessionsWithParams(IEnumerable<string> conditions, Dictionary<string, object> parameters, Database database)
    {
        var cmd = (await database.Select())
            .AddFields(["Id", "AccountId", "Token", "CreatedAt", "ExpiresAt"])
            .Table("Sessions")
            .AddConditions(conditions)
            .SetParameters(parameters)
            .Build();

        using var reader = await cmd.ExecuteReaderAsync();
        while (await reader.ReadAsync())
        {
            yield return new Session()
            {
                Id = reader.GetGuid(0),
                AccountId = reader.GetGuid(1),
                Token = reader.GetString(2),
                CreatedAt = reader.GetDateTime(3),
                ExpiresAt = reader.GetDateTime(4)
            };
        }
    }

    private static async IAsyncEnumerable<Login> GetLoginsWithParams(IEnumerable<string> conditions, Dictionary<string, object> parameters, Database database)
    {
        var cmd = (await database.Select())
            .AddFields(["Id", "AccountId", "CreatedAt", "IpAddress"])
            .Table("Logins")
            .AddConditions(conditions)
            .SetParameters(parameters)
            .Build();

        using var reader = await cmd.ExecuteReaderAsync();
        while (await reader.ReadAsync())
        {
            yield return new Login()
            {
                Id = reader.GetGuid(0),
                AccountId = reader.GetGuid(1),
                CreatedAt = reader.GetDateTime(2),
                IpAddress = reader.GetString(3)
            };
        }
    }
}

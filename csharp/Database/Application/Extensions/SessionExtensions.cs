using Dapper;
using Database.Application.Models;

namespace Database.Application.Extensions;

public static class SessionExtensions
{
    public static async Task<Session?> GetSessionById(this ApplicationDatabase database, Guid id) => 
        await (await database.GetConnectionAsync())
            .QueryFirstOrDefaultAsync<Session>(@$"
                SELECT * 
                FROM `Sessions`
                WHERE `Id`=@id
                LIMIT 1
            ", new {id});

    public static async Task<Session?> GetSessionByAccountId(this ApplicationDatabase database, Guid accountId) => 
        await (await database.GetConnectionAsync())
            .QueryFirstOrDefaultAsync<Session>(@$"
                SELECT * 
                FROM `Sessions`
                WHERE `AccountId`=@accountId
                LIMIT 1
            ", new {accountId});

    public static async Task<Session?> GetSessionByToken(this ApplicationDatabase database, string token) => 
        await (await database.GetConnectionAsync())
            .QueryFirstOrDefaultAsync<Session>(@$"
                SELECT * 
                FROM `Sessions`
                WHERE `Token`=@token
                LIMIT 1
            ", new {token});

    public static async Task<Session> CreateSession(this ApplicationDatabase database, Guid accountId)
    {
        if (!database.InTrasaction) 
            return await database.ExecuteInTransaction(
                async (_, _) => 
                    await database.CreateSession(accountId));

        var connection = await database.GetConnectionAsync();
        var session = await database.GetSessionByAccountId(accountId);
        if (session is not null && (await database.DeleteSession(accountId) == 0)) 
            throw new Exception($"Error deleting current account session for AccountId: {accountId}");    

        var sid = Guid.NewGuid();
        var token = database.GenerateJwtToken(sid, accountId);
        var createdAt = DateTime.UtcNow;
        var expiresAt = createdAt + TimeSpan.FromDays(7);

        var insertResult =  await connection.ExecuteAsync(@"
            INSERT INTO `Sessions`
                (`Id`, `AccountId`, `Token`, `CreatedAt`, `ExpiresAt`)
            VALUES
                (@id, @accountId, @token, @createdAt, @expiresAt)
        ", new {id=sid, accountId, token, createdAt, expiresAt});
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
        await (await database.GetConnectionAsync())
            .ExecuteAsync(@"
                DELETE FROM `Sessions`
                WHERE `AccountId`=@accountId
            ", new{accountId});
}

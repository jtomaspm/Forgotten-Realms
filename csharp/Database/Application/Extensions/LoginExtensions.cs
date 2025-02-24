using Dapper;
using Database.Application.Models;

namespace Database.Application.Extensions;

public static class LoginExtensions
{
    public static async Task<Login?> GetLoginById(this ApplicationDatabase database, Guid loginId) => 
        await (await database.GetConnectionAsync())
            .QueryFirstOrDefaultAsync<Login>(@$"
                SELECT * 
                FROM `Logins`
                WHERE `Id`=@id
                LIMIT 1
            ", new {id=loginId});

    public static async Task<IEnumerable<Login>> GetLoginByIpAddress(this ApplicationDatabase database, string ipAddress) => 
        await (await database.GetConnectionAsync())
            .QueryAsync<Login>(@$"
                SELECT * 
                FROM `Logins`
                WHERE `IpAddress`=@ipAddress
            ", new {ipAddress});

    public static async Task<Login> CreateLogin(this ApplicationDatabase database, Guid accountId, string ipAddress)
    {
        var connection = await database.GetConnectionAsync();
        var id = Guid.NewGuid();
        var createdAt = DateTime.UtcNow;

        var insertResult = await connection.ExecuteAsync(@"
            INSERT INTO `Logins`
                (`Id`, `AccountId`, `CreatedAt`, `IpAddress`)
            VALUES
                (@id, @accountId, @createdAt, @IpAddress)
        ", new{id, accountId, createdAt, ipAddress});
        if (insertResult != 1)
            throw new Exception("Error inserting new login record in database.");

        return new() { Id=id, AccountId = accountId, CreatedAt = createdAt, IpAddress = ipAddress };
    }
}
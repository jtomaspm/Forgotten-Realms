using System.IdentityModel.Tokens.Jwt;
using System.Reflection.Metadata.Ecma335;
using System.Security.Claims;
using System.Text;
using System.Threading.Tasks;
using Database.Application.Extensions;
using Database.Application.Models;
using Microsoft.IdentityModel.Tokens;

namespace Database.Application;

public class ApplicationDatabase : Database
{
    public ApplicationDatabase() : base(){}
    public ApplicationDatabase(DatabaseConfig config) : base(config){}
    internal string GenerateJwtToken(Guid sessionId, Guid accountId)
    {
        var securityKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(GetConfig().JwtSecret));
        var credentials = new SigningCredentials(securityKey, SecurityAlgorithms.HmacSha256);
        
        var claims = new[]
        {
            new Claim("sessionId", sessionId.ToString()),
            new Claim("accountId", accountId.ToString())
        };
        
        var token = new JwtSecurityToken(
            issuer: GetConfig().JwtIssuer,
            claims: claims,
            expires: DateTime.UtcNow.AddDays(7),
            signingCredentials: credentials
        );

        return new JwtSecurityTokenHandler().WriteToken(token);
    }

    public async Task<TokenValidationResponse?> ValidateJwtToken(string jwtToken) 
    {
        try 
        {
            var tokenHandler = new JwtSecurityTokenHandler();
            var key = Encoding.UTF8.GetBytes(GetConfig().JwtSecret);

            var validationParameters = new TokenValidationParameters
            {
                ValidateIssuer = true,
                ValidIssuer = GetConfig().JwtIssuer,
                ValidateAudience = false,
                ValidateIssuerSigningKey = true,
                IssuerSigningKey = new SymmetricSecurityKey(key),
                ValidateLifetime = true,
                ClockSkew = TimeSpan.Zero
            };

            var principal = tokenHandler.ValidateToken(jwtToken, validationParameters, out SecurityToken validatedToken);
            var claimsDict = principal.Claims.ToDictionary(c => c.Type, c => c.Value);

            if (claimsDict.TryGetValue("accountId", out string? accountId) && accountId is not null)
            {
                if(!Guid.TryParse(accountId, out var accountIdToken)) return null;
                var account = await this.GetAccountById(accountIdToken);
                if (account is null) return null;
                var result = new TokenValidationResponse() 
                {
                    Account = account,
                };
                if (claimsDict.TryGetValue("sessionId", out var sessionId))
                {
                    if(!Guid.TryParse(sessionId, out var sessionIdToken)) return null;
                    result.Session = await this.GetSessionById(sessionIdToken);
                    return result;
                }
            }
        }
        catch {return null;}
        return null;
    }
}


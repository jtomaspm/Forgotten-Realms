using System.Security.Cryptography;
using System.Text;

namespace Database.ApplicationDatabase.Models;

public class Session
{
    public Guid Id;
    public Guid AccountId;
    public required string Token;
    public DateTime CreatedAt;
    public DateTime ExpiresAt;

    public static string GenerateToken() =>
        Convert.ToBase64String(
            SHA256.HashData(
                Encoding.UTF8.GetBytes(
                    Guid.NewGuid().ToString())));
}




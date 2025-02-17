namespace Database.ApplicationDatabase.Models;

public class Session
{
    public Guid Id;
    public Guid AccountId;
    public required string Token;
    public DateTime CreatedAt;
    public DateTime ExpiresAt;
}




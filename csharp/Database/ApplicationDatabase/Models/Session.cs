namespace Database.ApplicationDatabase.Models;

public struct Session
{
    public Guid Id;
    public Guid AccountId;
    public string Token;
    public DateTime CreatedAt;
    public DateTime ExpiresAt;
}




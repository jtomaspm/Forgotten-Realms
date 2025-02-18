namespace Database.ApplicationDatabase.Models;

public class Account
{
    public Guid Id;
    public string? ExternalId;
    public string? Source;
    public required string Name;
    public required string Email;
    public required Role Role;
    public DateTime CreatedAt = DateTime.UtcNow;
    public DateTime UpdatedAt = DateTime.UtcNow;
}

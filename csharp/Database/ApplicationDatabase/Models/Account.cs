namespace Database.ApplicationDatabase.Models;

public class Account
{
    public Guid Id;
    public required string ExternalId;
    public required string Name;
    public required string Email;
    public required Role Role;
    public DateTime CreatedAt;
    public DateTime UpdatedAt;
}

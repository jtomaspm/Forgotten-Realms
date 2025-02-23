namespace Database.Application.Models;

public class Account
{
    public Guid Id;
    public string? ExternalId;
    public string? Source;
    public string? Name;
    public required string Email;
    public required Role Role;
    public DateTime CreatedAt;
    public DateTime UpdatedAt;
}

namespace Database.ApplicationDatabase.Models;

public struct Account
{
    public Guid Id;
    public string ExternalId;
    public string Name;
    public string Email;
    public Role Role;
    public DateTime CreatedAt;
    public DateTime UpdatedAt;
}

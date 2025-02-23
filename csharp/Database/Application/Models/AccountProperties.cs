namespace Database.Application.Models;

public class AccountProperties
{
    public Guid AccountId;
    public Guid VerificationToken;
    public DateTime TokenExpiresAt;
    public bool EmailVerified;
    public bool SendEmailNotifications;
    public DateTime CreatedAt;
    public DateTime UpdatedAt;
}
namespace Database.Application.Models;

public class TokenValidationResponse
{
    public required Account Account { get; set; }
    public Session? Session { get; set; }
}

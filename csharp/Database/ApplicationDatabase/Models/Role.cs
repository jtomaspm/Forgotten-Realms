
namespace Database.ApplicationDatabase.Models;

public enum RoleEnum 
{
    ADMIN,
    MODERATOR,
    PLAYER,
    GUEST
}

public class Role
{
    public RoleEnum Name;
}

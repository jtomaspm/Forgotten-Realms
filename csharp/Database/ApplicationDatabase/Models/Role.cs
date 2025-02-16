
namespace Database.ApplicationDatabase.Models;

public enum RoleEnum 
{
    ADMIN,
    MODERATOR,
    PLAYER,
    GUEST
}

public struct Role
{
    public RoleEnum Name;
}

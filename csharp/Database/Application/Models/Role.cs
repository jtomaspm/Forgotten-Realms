
namespace Database.Application.Models;

public enum RoleEnum 
{
    ADMIN,
    MODERATOR,
    NPC,
    PLAYER,
    GUEST
}

public class Role
{
    public required RoleEnum Name;

    public static Role Admin() => new(){Name=RoleEnum.ADMIN};
    public static Role Moderator() => new(){Name=RoleEnum.MODERATOR};
    public static Role Npc() => new(){Name=RoleEnum.NPC};
    public static Role Player() => new(){Name=RoleEnum.PLAYER};
    public static Role Guest() => new(){Name=RoleEnum.GUEST};
    public static Role[] AboveIncluding(RoleEnum roleName) => roleName switch
    {
        RoleEnum.ADMIN      => [Role.Admin()],
        RoleEnum.MODERATOR  => [Role.Moderator(), Role.Admin()],
        RoleEnum.NPC        => [Role.Npc(), Role.Moderator(), Role.Admin()],
        RoleEnum.PLAYER     => [Role.Player(), Role.Npc(), Role.Moderator(), Role.Admin()],
        RoleEnum.GUEST      => [Role.Guest(), Role.Player(), Role.Npc(), Role.Moderator(), Role.Admin()],
        _                   => [],
    };
    private static readonly Dictionary<string, Func<Role>> _roles = new () 
    {
        { "ADMIN", Role.Admin },
        { "MODERATOR", Role.Moderator },
        { "NPC", Role.Npc },
        { "PLAYER", Role.Player },
        { "GUEST", Role.Guest },
    };
    public static Role FromNameString(string name)
    {
        return _roles[name]();
    }
}

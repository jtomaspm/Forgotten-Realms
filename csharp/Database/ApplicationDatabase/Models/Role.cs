
namespace Database.ApplicationDatabase.Models;

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

    public static Role Player() => new(){Name=RoleEnum.PLAYER};
    public static Role Admin() => new(){Name=RoleEnum.ADMIN};
    public static Role Npc() => new(){Name=RoleEnum.NPC};
    public static Role Moderator() => new(){Name=RoleEnum.MODERATOR};
    public static Role Guest() => new(){Name=RoleEnum.GUEST};

    public Role[] AboveIncluding() => Name switch
    {
        RoleEnum.ADMIN      => [Role.Admin()],
        RoleEnum.MODERATOR  => [Role.Moderator(), Role.Admin()],
        RoleEnum.NPC        => [Role.Npc(), Role.Moderator(), Role.Admin()],
        RoleEnum.PLAYER     => [Role.Player(), Role.Npc(), Role.Moderator(), Role.Admin()],
        RoleEnum.GUEST      => [Role.Guest(), Role.Player(), Role.Npc(), Role.Moderator(), Role.Admin()],
        _                   => [],
    };
}

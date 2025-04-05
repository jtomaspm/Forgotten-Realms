export type UserRole = "admin" | "moderator" | "npc" | "player" | "guest"

export type User = {
    Id: string,
    Name: string,
    Email: string,
    Role: UserRole,
    Token: string
};
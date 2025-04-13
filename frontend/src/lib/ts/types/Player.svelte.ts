export type PlayerFaction = "caldari" | "varnak" | "dawnhold"
export type NpcFaction = "forgotten"
export type Faction = PlayerFaction | NpcFaction

export type Player = {
    id: string
    faction: PlayerFaction
    created_at: string
    updated_at: string
}
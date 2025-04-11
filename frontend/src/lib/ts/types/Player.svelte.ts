export type Faction = "caldari" | "varnak" | "dawnhold"

export type Player = {
    id: string
    faction: Faction
    created_at: string
    updated_at: string
}
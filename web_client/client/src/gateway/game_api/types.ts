export type Faction = "varnak" | "dawnhold" | "caldari"

export type Coords = {
    x: number,
    y: number,
}

export type Village = {
    coords: Coords,
    player: {id: number, name: string, faction: Faction},
    ally: {id: number, name: string},
    points: number,
}

export type Chunk = {
    position: Coords,
    villages: Map<string, Village>,
}
import type { Chunk, Coords, Village } from "./types";

export function coord_to_string(coord: Coords): string {
        return `${coord.x}|${coord.y}`;
}

export function get_chunk(coords: Coords, chunk_size: number): Chunk {
    const villages: Map<string, Village> = new Map();
    for (let x = coords.x; x < coords.x+chunk_size; x++) {
        for (let y = coords.y; y < coords.y+chunk_size; y++) {
            const seed = Math.random();
            if (Math.random() < 0.3) {
                const pid = Math.floor(seed * 100);
                const aid = Math.floor(seed * 10);
                villages.set(coord_to_string({x, y}), {
                    coords: {x, y},
                    points: Math.floor(Math.random() * 10000),
                    player: {id: pid, name: `player_${pid}`, faction: seed < 0.34 ? "caldari" : seed < 0.67 ? "dawnhold" : "varnak" },
                    ally: {id: aid, name: `ally_${aid}`}
                });
            }
        }
    }
    return {
        position: coords,
        villages: villages
    };
} 
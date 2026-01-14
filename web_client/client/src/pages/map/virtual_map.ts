import { coord_to_string, get_chunk } from "../../gateway/game_api/map";
import type { Chunk, Coords, Village } from "../../gateway/game_api/types";

export class VirtualMap {
    private _chunk_cache: Map<string, Chunk>;
    private readonly _chuck_size: number = 20;
    private readonly _map_size: number = 1000;

    constructor() {
        this._chunk_cache = new Map();
    }


    get_chunk_position(coords: Coords): Coords {
        return {
            x: coords.x - (coords.x % this._chuck_size),
            y: coords.y - (coords.y % this._chuck_size),
        }
    }

    get_village(coords: Coords): Village | undefined {
        //while (this._chunk_cache.size > 10) this._chunk_cache.delete(this._chunk_cache.keys().next().value!);
        const chunk_position = this.get_chunk_position(coords);
        let result = this._chunk_cache.get(coord_to_string(chunk_position));
        if (result) return result.villages.get(coord_to_string(coords)); 

        const new_chunk = get_chunk(chunk_position, this._chuck_size);
        this._chunk_cache.set(coord_to_string(chunk_position), new_chunk);

        return new_chunk.villages.get(coord_to_string(coords));
    }
}
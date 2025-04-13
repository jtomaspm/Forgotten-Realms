import type { PlayerFaction } from "../types/Player.svelte";

export class FactionState {
    faction: PlayerFaction = $state("caldari")

    SetFaction(f: PlayerFaction){
        this.faction = f;
    }
}
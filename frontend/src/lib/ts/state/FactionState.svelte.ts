import type { Faction } from "../types/Player.svelte";

export class FactionState {
    faction: Faction = $state("caldari")

    SetFaction(f: Faction){
        this.faction = f;
    }
}
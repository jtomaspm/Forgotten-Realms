<script lang="ts">
	import { Button, Card, Tabs } from "flowbite-svelte";
	import { AdjustmentsVerticalSolid } from "flowbite-svelte-icons";
	import Faction from "./Faction.svelte";
	import Caldari from "./factions/Caldari.svelte";
	import Varnak from "./factions/Varnak.svelte";
	import Dawnhold from "./factions/Dawnhold.svelte";
	import type { FactionState } from "$lib/ts/state/FactionState.svelte";
	import type { ErrorState } from "$lib/ts/state/ErrorState.svelte";

    let { registerCallback, faction, error } :
        { registerCallback: () => void, faction: FactionState, error: ErrorState } = $props();
</script>

<div class="flex items-center justify-center min-h-screen">
    <Card size="lg" padding="xl">
        <Tabs tabStyle="underline">
            <Faction {faction} name={"caldari"} Icon={AdjustmentsVerticalSolid} Content={Caldari} />
            <Faction {faction} name={"varnak"} Icon={AdjustmentsVerticalSolid} Content={Varnak} />
            <Faction {faction} name={"dawnhold"} Icon={AdjustmentsVerticalSolid} Content={Dawnhold} />
        </Tabs>
        <Button onclick={registerCallback} >Select</Button>
    </Card>
</div>


<style lang="postcss">
    @reference "tailwindcss";
</style>


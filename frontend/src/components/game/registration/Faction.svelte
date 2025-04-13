<script lang="ts">
	import type { FactionState } from "$lib/ts/state/FactionState.svelte";
	import type { PlayerFaction } from "$lib/ts/types/Player.svelte";
	import { TabItem } from "flowbite-svelte";
	import type { Component } from "svelte";

    let {
        name,
        Icon,
        Content, 
        faction
    }: {
        name: PlayerFaction
        Icon: Component
        Content: Component
        faction: FactionState
    } = $props();

</script>

<TabItem open={name===faction.faction} onclick={()=>faction.SetFaction(name)}>
    <div slot="title" class="tabtitle flex items-center gap-2">
        <Icon />
        {name}
    </div>
        <Content />
</TabItem>

<style lang="postcss">
    @reference "tailwindcss";
    .tabtitle {
        text-transform: capitalize;
    }
</style>
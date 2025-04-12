<script lang="ts">
	import type { RealmListing } from "$lib/ts/types/Realm.svelte";
	import { Button } from "flowbite-svelte";

    let { realm } :
        { realm: RealmListing } = $props();

    function onclick() {
        window.location.href = `/game/${realm.name}`
    }
</script>

<Button {onclick}>
    Play
</Button>

<style lang="postcss">
    @reference "tailwindcss";
</style>

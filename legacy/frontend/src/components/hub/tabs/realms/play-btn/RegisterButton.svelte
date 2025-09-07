<script lang="ts">
	import type { RealmListing } from "$lib/ts/types/Realm.svelte";
	import { Button } from "flowbite-svelte";

    let { realm } :
        { realm: RealmListing } = $props();

    function onclick() {
        window.location.href = `/callback/game?realm=${realm.id}`
    }
</script>

<Button {onclick}>
    Register
</Button>

<style lang="postcss">
    @reference "tailwindcss";
</style>


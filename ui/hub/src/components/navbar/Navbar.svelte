<script lang="ts">
	import type { User } from "$lib/ts/types/User.svelte";
	import Bottom from "./sections/Bottom.svelte";
	import Mid from "./sections/Mid.svelte";
	import Top from "./sections/Top.svelte";

    let { loggedIn }
        : { loggedIn: boolean } 
        = $props();
</script>

<div id="navbar">
    <div id="navbar-container">
        <Top />
        <Mid {loggedIn} />
        <Bottom />
    </div>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    #navbar {
        @apply w-[75px] h-screen bg-gray-900 fixed top-0 left-0 z-50;
    }
    #navbar-container {
        @apply h-full flex flex-col justify-between;
    }
</style>
<script lang="ts">
	import type { UserState } from "$lib/ts/state/UserState.svelte";
	import type { TabName } from "$lib/ts/types/NavbarTypes.svelte";
	import Bottom from "./sections/Bottom.svelte";
	import Mid from "./sections/Mid.svelte";
	import Top from "./sections/Top.svelte";

    let { activeTab, user, changeTab }
        : { activeTab: TabName, user: UserState, changeTab: (tab: TabName) => void } 
        = $props();
</script>

<div id="navbar">
    <div id="navbar-container">
        <Top />
        <Mid {activeTab} loggedIn={user.LoggedIn} {changeTab} />
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
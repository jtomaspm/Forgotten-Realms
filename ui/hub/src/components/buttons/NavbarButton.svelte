<script lang="ts">
	import type { TabName } from '$lib/ts/types/NavbarTypes.svelte';
    import { Button } from 'flowbite-svelte';
	import type { Component } from 'svelte';
    let { 
        onclick,
        tabName,
        Icon,
        activeTab,
    }:{
        onclick: () => void,
        tabName: TabName,
        Icon: Component,
        activeTab: TabName,
    } = $props()
    let active = $derived(activeTab===tabName)
</script>

<div class="container">
    <div class="panel left-pannel">
        <div class="active-indicator {active ? "active" : ""}"></div>
        <div class="dummy"></div>
    </div>
    <Button class="h-[48px] w-[48px] rounded-2xl hover:bg-gray-700/90 bg-transparent focus-within:ring-0 {active ? "bg-gray-700/90" : ""}" color="dark" {onclick}>
        <Icon class="scale-140" />
    </Button>
    <div class="pannel right-pannel"></div>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    @theme {
       --shadow-3xl: 0px 0px 70px 15px rgba(0.25, 0.25, 0.25, 0.25); 
    }
    .container {
        @apply transition-all duration-300 ease-in-out flex flex-row;
    }
    .pannel {
        @apply w-full h-full;
    }
    .left-pannel {
        @apply flex flex-row w-full;
    }
    .active-indicator {
        @apply rounded-r-xl flex-2 mt-[11px] mb-[11px];
    }
    .dummy {
        @apply flex-6;
    }
    .active {
        @apply shadow-3xl shadow-orange-700 bg-orange-700;
    }
</style>
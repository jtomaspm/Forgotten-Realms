<script lang="ts">
    import { ChartMixedDollarSolid, HomeSolid, GlobeSolid, UsersSolid, InfoCircleSolid, ShoppingBagSolid } from 'flowbite-svelte-icons';
	import NavbarButton from '../../buttons/NavbarButton.svelte';
	import type { TabName } from '$lib/ts/types/NavbarTypes.svelte';

    let activeTab : TabName = $state.raw("Home")

    function changeTab(tabName: TabName) {
        activeTab = tabName
    }

    let { loggedIn }
        : { loggedIn: boolean } 
        = $props();
</script>

<div>
    <NavbarButton onclick={()=>{changeTab("Home")}} Icon={HomeSolid} tabName="Home" {activeTab}/>
    {#if loggedIn}
    <NavbarButton onclick={()=>{changeTab("Play")}} Icon={GlobeSolid} tabName="Play" {activeTab} />
    <NavbarButton onclick={()=>{changeTab("Market")}} Icon={ChartMixedDollarSolid} tabName="Market" {activeTab} />
    <NavbarButton onclick={()=>{changeTab("Inventory")}} Icon={ShoppingBagSolid} tabName="Inventory" {activeTab} />
    {/if}
    <NavbarButton onclick={()=>{changeTab("Social")}} Icon={UsersSolid} tabName="Social" {activeTab} />
    <NavbarButton onclick={()=>{changeTab("Info")}} Icon={InfoCircleSolid} tabName="Info" {activeTab} />
</div>

<style lang="postcss">
    @reference "tailwindcss";
    div {
        @apply flex flex-col items-center gap-3 sticky top-1/2 transform -translate-y-1/2;
    }
</style>
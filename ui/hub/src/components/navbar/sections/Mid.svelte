<script lang="ts">
    import { ChartMixedDollarSolid, HomeSolid, GlobeSolid, UsersSolid, InfoCircleSolid, ShoppingBagSolid } from 'flowbite-svelte-icons';
	import NavbarButton from '../../buttons/NavbarButton.svelte';
	import type { TabName } from '$lib/ts/types/NavbarTypes.svelte';

    let { activeTab, loggedIn, changeTab }
        : { activeTab: TabName, loggedIn: boolean, changeTab: (tab: TabName) => void } 
        = $props();
</script>

<div>
    <NavbarButton onclick={()=>{changeTab("Home")}} Icon={HomeSolid} tabName="Home" {activeTab}/>
    <NavbarButton onclick={()=>{changeTab("Realms")}} Icon={GlobeSolid} tabName="Realms" {activeTab} />
    {#if loggedIn}
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
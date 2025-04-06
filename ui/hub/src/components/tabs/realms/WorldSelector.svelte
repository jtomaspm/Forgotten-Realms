<script lang="ts">
	import type { RealmListing } from '$lib/ts/types/Realm.svelte';
	import { Button, Dropdown, DropdownDivider, DropdownItem } from 'flowbite-svelte';
	import { ChevronDownOutline } from 'flowbite-svelte-icons';
    let { realm, setRealm, playing_realms, active_realms, closed_realms }
        : { realm: RealmListing | undefined, setRealm: (realm: RealmListing) => void, playing_realms: RealmListing[], active_realms: RealmListing[], closed_realms: RealmListing[] } 
        = $props();


    let dropdown: boolean = $state.raw(false);
</script>

<div>
    <Button color="dark">{realm?.name ?? "Select Realm"}<ChevronDownOutline class="w-6 h-6 ms-2 text-white dark:text-white" /></Button>
    <Dropdown bind:open={dropdown} class="bg-gray-800 text-white rounded-lg p-2">
        {#each playing_realms as li_realm}
		    <DropdownItem id={li_realm.id} onclick={()=>{dropdown=false;setRealm(li_realm)}} class="hover:bg-gray-900 hover:scale-110 transition-transform duration-100">{li_realm.name}</DropdownItem>
        {/each}
        {#if playing_realms.length > 0 && active_realms.length > 0}
		    <DropdownDivider />
        {/if}
        {#each active_realms as li_realm}
		    <DropdownItem id={li_realm.id} onclick={()=>{dropdown=false;setRealm(li_realm)}} class="hover:bg-gray-900 hover:scale-110 transition-transform duration-100">{li_realm.name}</DropdownItem>
        {/each}
        {#if (playing_realms.length > 0 && closed_realms.length > 0 && active_realms.length === 0) || (closed_realms.length > 0 && active_realms.length > 0)}
		    <DropdownDivider />
        {/if}
        {#each closed_realms as li_realm}
		    <DropdownItem id={li_realm.id} onclick={()=>{dropdown=false;setRealm(li_realm)}} class="hover:bg-gray-900 hover:scale-110 transition-transform duration-100">{li_realm.name}</DropdownItem>
        {/each}
    </Dropdown> 
</div>

<style lang="postcss">
    @reference "tailwindcss";
</style>
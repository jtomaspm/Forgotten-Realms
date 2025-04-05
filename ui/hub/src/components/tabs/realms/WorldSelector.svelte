<script lang="ts">
	import type { User } from '$lib/ts/types/User.svelte';
	import { Button, Dropdown, DropdownDivider, DropdownItem } from 'flowbite-svelte';
	import { ChevronDownOutline } from 'flowbite-svelte-icons';

    let { user }
        : { user: User } 
        = $props();
</script>

<div class="btn-container">
    <Button color="dark">Dropdown button<ChevronDownOutline class="w-6 h-6 ms-2 text-white dark:text-white" /></Button>
    <Dropdown color="dark">
        <DropdownItem>Dashboard</DropdownItem>
        <DropdownDivider />
        <DropdownItem>Settings</DropdownItem>
        <DropdownItem>Earnings</DropdownItem>
        <DropdownItem slot="footer">Separated link</DropdownItem>
    </Dropdown> 
</div>

<style lang="postcss">
    @reference "tailwindcss";
    
    .btn-container {
        @apply fixed top-4 left-1/2 -translate-x-1/2 z-50;
    }
</style>
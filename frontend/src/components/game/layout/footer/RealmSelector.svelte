<script lang="ts">
	import type { PlayableRealm } from '$lib/ts/types/Realm.svelte';
    import { Button, Dropdown, DropdownItem } from 'flowbite-svelte';

    const { realm, realms } : { realm: PlayableRealm, realms: PlayableRealm[] } = $props();
    const choices = realms.filter(x => x.id !== realm.id);
</script>

<Button class="p-0 pl-[5px] pr-[5px] ml-[5px] mr-[5px] rounded-sm focus-within:ring-0" color="dark">
    {realm.name}
</Button>
<Dropdown>
    {#each choices as choice}
        <DropdownItem onclick={()=>{ window.location.href = `/game/${choice.name}`; }}>{choice.name}</DropdownItem>
    {/each}
</Dropdown>
<script lang="ts">
    import {  } from 'flowbite-svelte-icons';
	import type { User } from '$lib/ts/types/User.svelte';
    import { Button } from 'flowbite-svelte';

    let { user, loggedIn }
        : { user: User | undefined, loggedIn: boolean } 
        = $props();
</script>

<Button color="dark">
    Profile
</Button>

<style lang="postcss">
    @reference "tailwindcss";
</style>
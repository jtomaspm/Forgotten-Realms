<script lang="ts">
    import {  } from 'flowbite-svelte-icons';
	import type { User } from '$lib/ts/types/User.svelte';
	import ProfileButton from './ProfileButton.svelte';
	import LoginButton from './LoginButton.svelte';

    let { user, loggedIn }
        : { user: User | undefined, loggedIn: boolean } 
        = $props();
</script>

{#if loggedIn}
    <ProfileButton {user} {loggedIn} />
{:else}
    <LoginButton />
{/if}

<style lang="postcss">
    @reference "tailwindcss";
</style>


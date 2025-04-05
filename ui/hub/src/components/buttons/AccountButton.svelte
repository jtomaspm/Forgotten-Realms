<script lang="ts">
    import {  } from 'flowbite-svelte-icons';
	import type { User } from '$lib/ts/types/User.svelte';
	import ProfileButton from './ProfileButton.svelte';
	import LoginButton from './LoginButton.svelte';

    let { user, loggedIn }
        : { user: User | undefined, loggedIn: boolean } 
        = $props();
</script>

<div class="btn-container">
    {#if loggedIn}
        <ProfileButton {user} {loggedIn} />
    {:else}
        <LoginButton />
    {/if}
</div>

<style lang="postcss">
    @reference "tailwindcss";
    .btn-container {
        @apply fixed top-6 right-6 z-50;
    }
</style>


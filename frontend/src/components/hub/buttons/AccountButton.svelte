<script lang="ts">
    import {  } from 'flowbite-svelte-icons';
	import type { User } from '$lib/ts/types/User.svelte';
	import ProfileButton from './ProfileButton.svelte';
	import LoginButton from './LoginButton.svelte';
	import type { UserState } from '$lib/ts/state/UserState.svelte';

    let { user }
        : { user: UserState } 
        = $props();
</script>

<div class="btn-container">
    {#if user.LoggedIn}
        <ProfileButton />
    {:else}
        <LoginButton />
    {/if}
</div>

<style lang="postcss">
    @reference "tailwindcss";
    .btn-container {
        @apply fixed top-4 right-4 z-50;
    }
</style>


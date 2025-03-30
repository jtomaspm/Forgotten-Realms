<script lang="ts">
    import {  } from 'flowbite-svelte-icons';
	import type { User } from '$lib/ts/types/User.svelte';
    import { Button } from 'flowbite-svelte';

    let { user, loggedIn }
        : { user: User | undefined, loggedIn: boolean } 
        = $props();
</script>

<div class="profile-btn-container">
    <Button color="dark">
        Profile
    </Button>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    .profile-btn-container {
        @apply fixed top-6 right-6 z-50;
    }
</style>
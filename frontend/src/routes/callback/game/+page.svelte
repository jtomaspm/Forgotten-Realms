<script lang="ts">
	import CreatePlayer from "../../../components/game/registration/CreatePlayer.svelte";
	import type { AuthLoginCallbackResponse, AuthRegistrationCallbackResponse } from '$lib/ts/types/AuthCallbackResponse.svelte';
    import { onMount } from 'svelte';
    const authUrl = import.meta.env.VITE_AUTH_URL;


    onMount(async () => {
        const urlParams = new URLSearchParams(window.location.search);
        const realmId = urlParams.get('realm');
        
        if (!realmId) {
            window.location.href = "/";
            return;
        }

    });
</script>

<div>
    <CreatePlayer />
</div>
<script lang="ts">
	import type { AuthCallbackResponse } from '$lib/ts/types/AuthCallbackResponse';
    import { onMount } from 'svelte';
	import CreateAccount from '../../../components/registration/CreateAccount.svelte';
    const authUrl = import.meta.env.VITE_AUTH_URL;

    let errorMessage = $state.raw("");
    let createUser : AuthCallbackResponse | undefined = $state.raw();

    onMount(async () => {
        const urlParams = new URLSearchParams(window.location.search);
        const code = urlParams.get('code');
        
        if (!code) {
            errorMessage = "Missing code parameter.";
            return;
        }

        try {
            const response = await fetch(`${authUrl}/api/github/callback?code=${code}`, {
                method: 'GET',
            });
            
            if (response.status === 200) {
                const data = await response.json();
                console.log('Login successful:', data);
                return;
            }

            if (response.status === 405) {
                createUser = await response.json();
                return;
            }

            const errorData = await response.json();
            errorMessage = errorData.error || "Failed to login";

        } catch (error) {
            errorMessage = 'An error occurred while processing the login.';
        }
    });
</script>

<main>
  {#if errorMessage}
    <p style="color: red;">{errorMessage}</p>
  {/if}
  {#if createUser && !errorMessage}
    <CreateAccount {createUser} />
  {/if}
</main>
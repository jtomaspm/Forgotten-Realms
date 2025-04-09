<script lang="ts">
	import type { AuthLoginCallbackResponse, AuthRegistrationCallbackResponse } from '$lib/ts/types/AuthCallbackResponse.svelte';
    import { onMount } from 'svelte';
	import CreateAccount from '../../../components/hub/registration/CreateAccount.svelte';
    const authUrl = import.meta.env.VITE_AUTH_URL;

    let errorMessage = $state.raw("");
    let authResponse : AuthRegistrationCallbackResponse | undefined = $state.raw();

    onMount(async () => {
        const urlParams = new URLSearchParams(window.location.search);
        const code = urlParams.get('code');
        
        if (!code) {
            window.location.href = "/";
            return;
        }

        try {
            const response = await fetch(`${authUrl}/api/github/callback?code=${code}`, {
                method: 'GET',
            });
            
            if (response.status === 200) {
                const data: AuthLoginCallbackResponse = await response.json();
                localStorage.setItem("popfrsid", data.token);
                window.location.href = "/";
                return;
            }

            if (response.status === 405) {
                authResponse = await response.json();
                return;
            }

            const errorData = await response.json();
            errorMessage = errorData.error || "Failed to login";

        } catch (error) {
            errorMessage = 'An error occurred while processing the login.';
        }
    });
</script>

<div>
  {#if errorMessage}
    <p style="color: red;">{errorMessage}</p>
  {/if}
  {#if authResponse && !errorMessage}
    <CreateAccount {authResponse} />
  {/if}
</div>
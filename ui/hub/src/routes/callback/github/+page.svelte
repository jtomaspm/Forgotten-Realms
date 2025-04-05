<script>
  import { onMount } from 'svelte';

  let errorMessage = '';

  onMount(async () => {
    const urlParams = new URLSearchParams(window.location.search);
    const code = urlParams.get('code');
    
    if (!code) {
        errorMessage = "Missing code parameter.";
        return;
    }

    try {
        const response = await fetch(`http://localhost:8080/api/github/callback?code=${code}`, {
            method: 'GET',
        });
        
        if (response.status === 200) {
            const data = await response.json();
            console.log('Login successful:', data);
            return;
        }

        if (response.status === 405) {
            const data = await response.json();
            console.log('Create  account:', data);
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
</main>
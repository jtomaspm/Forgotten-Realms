<script lang="ts">
	import type { AccountCreated, AuthRegistrationCallbackResponse } from "$lib/ts/types/AuthCallbackResponse";
	import { Button, Modal } from "flowbite-svelte";
	import { ArrowLeftOutline, CheckCircleOutline, CheckCircleSolid, ExclamationCircleOutline } from "flowbite-svelte-icons";
    const authUrl = import.meta.env.VITE_AUTH_URL;
    let { authResponse, accountCreated }
        : { authResponse: AuthRegistrationCallbackResponse, accountCreated: AccountCreated }
        = $props();

    let modal = $state.raw(false)
    let error = $state.raw(false)
    let modalTxt = $state.raw("")

    function returnHome() {
        window.location.href = '/';
    }
    async function verifyEmail() {
        try {
            const response = await fetch(`${authUrl}/api/account/verify?token=${accountCreated.token}`, {
                method: 'GET',
            });
            if (response.status === 202) {
                error = false;
                modalTxt = "Email verified.";
                accountCreated.token = "";
                modal = true;
            }
        } catch (ex) {
            console.log(ex);
            error = true;
            modalTxt = "Failed to verify email.";
            modal = true;
        }
    }
</script>


<Modal bind:open={modal} size="xs" autoclose outsideclose>
  <div class="text-center">
    {#if error}
        <ExclamationCircleOutline class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200" />
    {:else} 
        <CheckCircleOutline class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200" />
    {/if}
    <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">{modalTxt}</h3>
  </div>
</Modal>

<div>
    <h5>Welcome {accountCreated.name}</h5>
    <p>Please verify your account using the link sent to {authResponse.email}</p>
    
    <div class="buttons-container">
    <Button onclick={returnHome} class="w-fit">
        <ArrowLeftOutline class="w-6 h-6 text-white mr-[5px]" /> Home 
    </Button>

    {#if !(accountCreated.token === "")}
        <Button onclick={verifyEmail} class="w-fit">
            <CheckCircleSolid class="w-6 h-6 text-white mr-[5px]" /> Verify
        </Button>
    {/if}
    </div>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    .buttons-container {
        @apply mt-[10px];
    }
    h5 {
        @apply mb-2 text-2xl font-bold tracking-tight text-gray-900;
    }
    p {
        @apply mb-3 font-normal text-gray-700 leading-tight;
    }
</style>

<script lang="ts">
	import type { AccountCreated, AuthRegistrationCallbackResponse } from "$lib/ts/types/AuthCallbackResponse.svelte";
	import { Button, Modal } from "flowbite-svelte";
	import { ArrowRightOutline, ExclamationCircleOutline } from "flowbite-svelte-icons";
    const authUrl = import.meta.env.VITE_AUTH_URL;
    
    let errorMessage = $state.raw("")
    let error = $state(false)
    $effect(()=>{
        if (error  === false)  {
            errorMessage = ""
        }
    });
    let name = $state.raw("")
    let sendEmailNotifications =  $state.raw(false)
    let acceptTerms =  $state.raw(false)
    let { authResponse, accountCreated }
        : { authResponse: AuthRegistrationCallbackResponse, accountCreated: AccountCreated }
        = $props();
    function submitName()  {
        accountCreated.created = false;
        if (name.trim() === "") {
            errorMessage = "Name cannot be empty."
            error = true;
            return;
        }
        if (!acceptTerms) {
            errorMessage = "Must accept terms and conditions."
            error = true;
            return;
        }
        fetch(`${authUrl}/api/account`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': "Bearer " +  authResponse.token
            },
            body: JSON.stringify({
                name: name,
                send_email_notifications: sendEmailNotifications
            })
        }).then(response => {
            if (response.status === 201) {
                response.json().then(data => {
                    accountCreated.token = data.token ?? "";
                    accountCreated.name = name;
                    accountCreated.created = true;
                }).catch(error => {
                    errorMessage = "Error handling api response."
                    error = true;
                });
                return;
            }
            response.json().then(errorData => {
                errorMessage = errorData.error || "Failed to login";
                error = true;
            });
        }).catch(error => {
            errorMessage = 'An error occurred while processing the login.';
            error = true;
        });
    }
</script>

<Modal bind:open={error} size="xs" autoclose outsideclose>
  <div class="text-center">
    <ExclamationCircleOutline class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200" />
    <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">{errorMessage}</h3>
  </div>
</Modal>

<div>
    <h5>Welcome to the Forgotten Realms</h5>
    <p>What is your name?</p>
    
    <input type="text" name="accountName" id="accountName" bind:value={name} />
    <div class="flex items-center">
        <input type="checkbox" id="sendEmailNotifications" bind:checked={sendEmailNotifications} />
        <label for="sendEmailNotifications">Receive Email Notifications</label>
    </div>
    <div class="flex items-center">
        <input type="checkbox" id="acceptTerms" bind:checked={acceptTerms} />
        <label for="acceptTerms">
            Accept <a href="https://github.com/jtomaspm/Forgotten-Realms/blob/main/CODE_OF_CONDUCT.md" target="_blank">Terms and Conditions</a>
        </label>
    </div>
    
    <Button onclick={submitName} class="w-fit mt-[10px]">
        Create Account <ArrowRightOutline class="w-6 h-6 ms-2 text-white" />
    </Button>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    #accountName  {
        @apply mb-[10px];
    }
    h5 {
        @apply mb-2 text-2xl font-bold tracking-tight text-gray-900;
    }
    p {
        @apply mb-3 font-normal text-gray-700 leading-tight;
    }
    label {
        @apply ml-2 text-gray-700;
    }
    a {
        @apply underline;
    }
</style>

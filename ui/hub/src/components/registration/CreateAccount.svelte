<script lang="ts">
	import type { AccountCreated, AuthRegistrationCallbackResponse } from "$lib/ts/types/AuthCallbackResponse";
	import { Card } from "flowbite-svelte";
	import NameForm from "./NameForm.svelte";
	import VerifyEmail from "./VerifyEmail.svelte";
    let { authResponse }
        : { authResponse: AuthRegistrationCallbackResponse } 
        = $props();
    let accountCreated: AccountCreated = $state({
        token: "",
        created: false,
        name: "",
    })
</script>

<div class="flex items-center justify-center min-h-screen">
    <Card size="md" padding="xl">
        {#if !accountCreated.created}
            <NameForm {authResponse} {accountCreated} />
        {:else}
            <VerifyEmail {authResponse} {accountCreated} />
        {/if}
    </Card>
</div>


<style lang="postcss">
    @reference "tailwindcss";
</style>

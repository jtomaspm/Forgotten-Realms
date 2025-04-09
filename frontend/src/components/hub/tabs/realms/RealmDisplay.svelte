<script lang="ts">
	import type { RealmListing } from "$lib/ts/types/Realm.svelte";
	import type { User } from "$lib/ts/types/User.svelte";
	import PlayButton from "./play-btn/PlayButton.svelte";
	import RegisterButton from "./play-btn/RegisterButton.svelte";
    import { DateTime } from "ts-luxon";

    let { user, loggedIn, realm }
        : { user: User | undefined, loggedIn: boolean, realm: RealmListing } 
        = $props();
    let openDays = $derived(Math.abs(Math.trunc(DateTime.fromISO(realm.created_at, { zone: 'utc' }).diffNow(['days']).days)));
</script>

<div class="display-wrapper">
    <div class="display-container">
        <div class="title-container">
            <h1>{realm.name}</h1>
            <p>Started {openDays === 0 ? "today" : openDays === 1 ? "yesterday" : `${openDays} days ago`}</p>
        </div>
        {#if  !loggedIn}
            <p>Must be signed in to play in the forgotten realms.</p>
        {:else if realm.registered && realm.status !== 'ended'}
            <div class="btn-container">
                <PlayButton />
            </div>
        {:else if !realm.registered && realm.status === 'open'}
            <div class="btn-container">
                <RegisterButton {realm} />
            </div>
        {:else if realm.status === 'ended'}
            <div class="realm-status">
                <p>Realm has ended. You can no longer play on it.</p>
            </div>
        {:else if realm.status === 'closed'}
            <div class="realm-status">
                <p>Realm is closed. No new players allowed.</p>
            </div>
        {:else}
            <div class="realm-status">
                <p>Can't access realm for unknown reasons.</p>
            </div>
        {/if}
    </div>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    .display-wrapper {
        @apply pt-[80px] pl-[175px] pr-[100px] min-w-[500px] w-auto;
    }
    .display-container {
        @apply w-auto;
    }
    .btn-container {
        @apply mb-[10px];
    }
    h1 {
        @apply text-3xl mr-[10px];
    }
    .title-container {
        @apply flex mb-[15px] items-end;
    }
    .title-container p {
        @apply text-gray-500;
    }
</style>

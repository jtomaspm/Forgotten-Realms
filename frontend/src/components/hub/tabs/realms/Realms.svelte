<script lang="ts">
	import type { RealmListing } from "$lib/ts/types/Realm.svelte";
	import { onMount } from "svelte";
	import RealmDisplay from "./RealmDisplay.svelte";
	import WorldSelector from "./WorldSelector.svelte";
	import { GetRealms } from "$lib/ts/sdk/hub/Realms.svelte";
	import type { UserState } from "$lib/ts/state/UserState.svelte";
    const hubUrl = import.meta.env.VITE_HUB_URL;

    let { user, realm, setRealm }
        : { user: UserState, realm: RealmListing | undefined, setRealm: (realm: RealmListing) => void } 
        = $props();
    let realms: RealmListing[]  = $state.raw([]);
    let playing_realms = $derived(
        realms
            .filter(r => r.registered && (r.status !== 'ended'))
            .sort((a, b) => Date.parse(b.created_at) - Date.parse(a.created_at))
    );
    let active_realms = $derived(
        realms
            .filter(r => !r.registered && r.status === 'open')
            .sort((a, b) => Date.parse(b.created_at) - Date.parse(a.created_at))
    );
    let closed_realms = $derived(
        realms
            .filter(r => !r.registered && r.status !== 'open')
            .sort((a, b) => Date.parse(b.created_at) - Date.parse(a.created_at))
    );

    onMount(async () => {
        let response = await GetRealms({url: hubUrl}, (user.LoggedIn) ? user.Get()!.Token : undefined)
        if (response.error) {
            return
        }
        realms = response.realms;
        if (realm) {
            return;
        }
        if (playing_realms.length > 0) {
            realm = playing_realms[0];
        }else if (active_realms.length > 0) {
            realm =  active_realms[0];
        }else if (closed_realms.length > 0) {
            realm = closed_realms[0]
        }
    });
</script>

<div class="top-btn-container">
    <WorldSelector {realm} {setRealm} {playing_realms} {active_realms} {closed_realms} />
</div>

{#if realm}
    <RealmDisplay {user} {realm} />
{/if}

<style lang="postcss">
    @reference "tailwindcss";
    
    .top-btn-container {
        @apply fixed top-4 left-1/2 -translate-x-1/2 z-50;
    }
</style>
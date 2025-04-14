<script lang="ts">
	import { GetAccount } from "$lib/ts/sdk/auth/Account.svelte";
	import { GetPlayer, RegisterPlayer } from "$lib/ts/sdk/game/Player.svelte";
	import { GetRealm } from "$lib/ts/sdk/hub/Realms.svelte";
	import { ErrorState } from "$lib/ts/state/ErrorState.svelte";
	import { FactionState } from "$lib/ts/state/FactionState.svelte";
	import { GetSessionToken } from "$lib/ts/store/Browser.svelte";
	import type { Realm } from "$lib/ts/types/Realm.svelte";
	import CreatePlayer from "../../../components/game/registration/CreatePlayer.svelte";
    import { onMount } from 'svelte';
    const authUrl = import.meta.env.VITE_AUTH_URL;
    const hubUrl = import.meta.env.VITE_HUB_URL;

    const faction: FactionState = new FactionState();
    const error: ErrorState = new ErrorState(); 

    let realm: Realm | undefined = $state.raw();
    let token : string = $state.raw("")

    async function registerCallback() {
        if (!realm) {
            return;
        }
        console.log(realm);
        const result = await RegisterPlayer({url: realm.api}, token, faction.faction, "random");
        if (result.error) {
            error.SetMessage(result.error.Errors[0]);
            error.Error();
            return;
        }
        console.log(result);
        window.location.href = `/game/${realm.name}/`
    }

    function exit() {
        window.location.href = "/";
    }

    onMount(async () => {
        let t = GetSessionToken();
        if (!t) {
            exit();
        }
        token = t!;
        const urlParams = new URLSearchParams(window.location.search);
        const realmId = urlParams.get('realm');
        if (!realmId) {
            exit();
            return;
        }
        const res = await GetRealm({url: hubUrl}, realmId);
        if (res.error || !res.realm) {
            exit();
            return;
        }
        realm = res.realm;
        let user = await GetAccount({url: authUrl}, token);
        if(user.error || !user.account) {
            exit();
        }
        let player = await GetPlayer({url: realm.api}, user.account!.Id)
        if(!player.error) {
            window.location.href = `/game/${realm.name}/`
        } 
    });
</script>

<div>
    <CreatePlayer {faction} {registerCallback} {error}/>
</div>
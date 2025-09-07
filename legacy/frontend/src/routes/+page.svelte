<script lang="ts">
	import type { User } from '$lib/ts/types/User.svelte';
	import { onMount } from 'svelte';
	import AccountButton from '../components/hub/buttons/AccountButton.svelte';
	import Navbar from '../components/hub/navbar/Navbar.svelte';
	import type { TabName } from '$lib/ts/types/NavbarTypes.svelte';
	import Home from '../components/hub/tabs/home/Home.svelte';
	import Realms from '../components/hub/tabs/realms/Realms.svelte';
	import Inventory from '../components/hub/tabs/inventory/Inventory.svelte';
	import Market from '../components/hub/tabs/market/Market.svelte';
	import Social from '../components/hub/tabs/social/Social.svelte';
	import Info from '../components/hub/tabs/info/Info.svelte';
	import type { RealmListing } from '$lib/ts/types/Realm.svelte';
	import { DeleteSessionToken, GetSessionToken } from '$lib/ts/store/Browser.svelte';
	import { GetAccount } from '$lib/ts/sdk/auth/Account.svelte';
	import { UserState } from '$lib/ts/state/UserState.svelte';
    const authUrl = import.meta.env.VITE_AUTH_URL;

    let activeTab: TabName = $state.raw("Home");
    function changeTab(tab: TabName) {
        activeTab = tab
    }
    let realm: RealmListing | undefined = $state.raw();
    const setRealm = (r: RealmListing) => {
        realm = r;
    };


	let user : UserState = new UserState();
    onMount(async () => {
        let token = GetSessionToken();
        if (!token) {
            return;
        }
        const response = await GetAccount({url:authUrl}, token)
        console.log(response);
        if (response.error === undefined && response.account !== null) {
            user.Set(response.account);
            return;
        }
        DeleteSessionToken();
    });
</script>

{#if activeTab === "Home"}
    <Home />
{:else if activeTab === "Realms"}
    <Realms {user} {realm} {setRealm} />
{:else if activeTab === "Market" && user}
    <Market />
{:else if activeTab === "Inventory" && user}
    <Inventory />
{:else if activeTab === "Social"}
    <Social />
{:else if activeTab === "Info"}
    <Info />
{/if}

<Navbar {activeTab} {user} {changeTab} />
<AccountButton {user} />
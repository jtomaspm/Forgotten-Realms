<script lang="ts">
	import type { User } from '$lib/ts/types/User.svelte';
	import { onMount } from 'svelte';
	import AccountButton from '../components/buttons/AccountButton.svelte';
	import Navbar from '../components/navbar/Navbar.svelte';
	import type { TabName } from '$lib/ts/types/NavbarTypes.svelte';
	import Home from '../components/tabs/home/Home.svelte';
	import Realms from '../components/tabs/realms/Realms.svelte';
	import Inventory from '../components/tabs/inventory/Inventory.svelte';
	import Market from '../components/tabs/market/Market.svelte';
	import Social from '../components/tabs/social/Social.svelte';
	import Info from '../components/tabs/info/Info.svelte';
	import type { RealmListing } from '$lib/ts/types/Realm.svelte';
    const authUrl = import.meta.env.VITE_AUTH_URL;

    let activeTab: TabName = $state.raw("Home");
    function changeTab(tab: TabName) {
        activeTab = tab
    }
    let realm: RealmListing | undefined = $state.raw();
    const setRealm = (r: RealmListing) => {
        realm = r;
    };


	let user : User | undefined = $state.raw()
    let loggedIn = $derived(user != undefined && !(user.Token === ""))
    onMount(async () => {
        let token = localStorage.getItem("popfrsid")
        if (!token) {
            return;
        }
        const response = await fetch(`${authUrl}/api/account?token=${token}`, {
            method: 'GET',
        });
        if (response.status === 200) {
            let data = await response.json();
            user = {
                Id: data.id,
                Email: data.email,
                Name: data.name,
                Role: data.role,
                Token: token
            };
            return;
        }
        localStorage.removeItem("popfrsid")
    });
</script>

{#if activeTab === "Home"}
    <Home />
{:else if activeTab === "Realms"}
    <Realms {user} {loggedIn} {realm} {setRealm} />
{:else if activeTab === "Market" && user}
    <Market />
{:else if activeTab === "Inventory" && user}
    <Inventory />
{:else if activeTab === "Social"}
    <Social />
{:else if activeTab === "Info"}
    <Info />
{/if}

<Navbar {activeTab} {loggedIn} {changeTab} />
<AccountButton {user} {loggedIn} />
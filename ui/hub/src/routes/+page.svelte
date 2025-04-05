<script lang="ts">
	import type { User } from '$lib/ts/types/User.svelte';
	import AccountButton from '../components/buttons/AccountButton.svelte';
	import Navbar from '../components/navbar/Navbar.svelte';

	let user : User | undefined = $state.raw()
    let loggedIn = $derived(user != undefined && user.Token != "")
</script>

<Navbar {loggedIn} />
<AccountButton {user} {loggedIn} />
<script lang="ts">
	import type { User } from '$lib/ts/types/User.svelte';
	import '../app.css';
	import ProfileButton from '../components/buttons/ProfileButton.svelte';
	import Navbar from '../components/navbar/Navbar.svelte';

	let { children } = $props();
	let user : User | undefined = $state.raw({
		Id: "06b19b30-297d-4b87-b0fc-89102c3c6a02",
		Name: "PopAndBoom",
		Email: "joaotomachado@proton.me",
		Token: "3286ffae-c40f-4d68-ad73-2c2f9a5d337d",
	})
    let loggedIn = $derived(user != undefined && user.Token != "")
</script>

<div id="background">
<Navbar {loggedIn} />
<ProfileButton {user} {loggedIn} />
{@render children()}
</div>

<style lang="postcss">
    @reference "tailwindcss";
    #background {
        @apply w-screen h-screen;
        background-image: url('$lib/static/background.jpg');
    }
</style>
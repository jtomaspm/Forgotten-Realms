<script lang="ts">
	import '../../../app.css';
	import Navbar from '../../../components/game/layout/navbar/Navbar.svelte'
	import Footer from '../../../components/game/layout/footer/Footer.svelte'
	import Screen from '../../../components/game/layout/screen/Screen.svelte';
	import { setContext } from 'svelte';
	import { redirect } from '@sveltejs/kit';
	let { data, children } = $props();
	if(!data.realm || !data.realms) {
		throw redirect(302, "/")
	}
	setContext('realm', data.realm);
	setContext('realms', data.realms);
</script>

<Navbar realm={data.realm} />
<Screen>
	{@render children()}
</Screen>
<Footer realm={data.realm} realms={data.realms} />

<style lang="postcss">
	@reference "tailwindcss";
</style>
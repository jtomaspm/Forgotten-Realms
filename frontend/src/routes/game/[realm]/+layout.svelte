<script lang="ts">
	import '../../../app.css';
	import Navbar from '../../../components/game/layout/navbar/Navbar.svelte'
	import Footer from '../../../components/game/layout/footer/Footer.svelte'
	import Screen from '../../../components/game/layout/screen/Screen.svelte';
	let { children } = $props();
</script>

<Navbar />
<Screen>
	{@render children()}
</Screen>
<Footer />

<style lang="postcss">
	@reference "tailwindcss";
</style>
<script lang="ts">
	import '../app.css';
	import Navbar from '../components/layout/navbar/Navbar.svelte'
	import Footer from '../components/layout/footer/Footer.svelte'
	import Screen from '../components/layout/screen/Screen.svelte';
	let { children } = $props();
</script>

<Navbar />
<Screen>
	{@render children()}
</Screen>
<Footer />

<style lang="postcss">
	@reference "tailwindcss";
    :global(html, body) {
        @apply h-full;
    }
</style>
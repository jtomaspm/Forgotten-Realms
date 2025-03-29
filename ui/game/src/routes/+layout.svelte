<script lang="ts">
	import '../app.css';
	import Navbar from '../components/layout/navbar/Navbar.svelte'
	import Footer from '../components/layout/footer/Footer.svelte'
	let { children } = $props();
</script>

<Navbar />
{@render children()}
<Footer />

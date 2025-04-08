<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<div id="background">
{@render children()}
</div>

<style lang="postcss">
    @reference "tailwindcss";
    #background {
        @apply w-screen min-h-screen;
        background-image: url('$lib/static/background.jpg');
        background-attachment: fixed;
        background-size: auto;
        background-repeat: no-repeat;
        background-position: top center;
    }
</style>
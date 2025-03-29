<script lang="ts">
    import ResourceBar from "./resource_bar/ResourceBar.svelte";
    let { children } = $props();
</script>

<div id="window-container">
    <ResourceBar />
    {@render children()}
</div>

<style lang="postcss">
    @reference "tailwindcss";
    #window-container {
        @apply w-full mb-[50px] bg-yellow-200;
    }
</style>
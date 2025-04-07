<script lang="ts">
	import Resources from "./resources/Resources.svelte";
</script>

<div id="resource-bar" class="flex flex-col">
    <div id="top-bar-cover"></div>
    <div id="bar-container">
        <div class="side-panel"></div>
        <Resources />
        <div class="side-panel"></div>
    </div>
    <div id="bottom-margin"></div>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    #top-bar-cover {
        @apply w-full h-[40px];
    }
    #bottom-margin {
        @apply w-full h-[10px];
    }
    #bar-container {
        @apply flex flex-row w-full h-[40px];
    }
    .side-panel {
        @apply flex-1 h-auto;
    }
</style>
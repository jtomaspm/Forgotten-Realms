<script lang="ts">
    let { children } = $props();
</script>

<div id="container">
    <div class="side-panel"></div>
    <div id="window-container">
        {@render children()}
    </div>
    <div class="side-panel"></div>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    #container {
        @apply flex;
    }
    #window-container {
        @apply w-[980px] p-[5px] mb-[10px] bg-yellow-500;
    }
    .side-panel {
        @apply flex-1 h-auto;
    }
</style>
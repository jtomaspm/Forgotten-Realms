<script lang="ts">
	import Resources from "./resources/Resources.svelte";
</script>

<div id="dummy"></div>
<div id="bar-container">
    <div class="side-panel"></div>
    <Resources />
    <div class="side-panel"></div>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    #dummy {
        @apply w-full h-[40px];
    }
    #bar-container {
        @apply flex flex-row w-full h-[50px] mb-[30px];
    }
    .side-panel {
        @apply flex-1 h-auto;
    }
</style>
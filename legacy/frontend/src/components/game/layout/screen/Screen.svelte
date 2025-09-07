<script lang="ts">
    import GameWindow from "./game_window/GameWindow.svelte";
	let { children } = $props();
</script>

<div id="screen-container">
    <div class="side-panel">
    </div>

    <div id="game-panel">
        <GameWindow>
            {@render children()}
        </GameWindow>
    </div>

    <div class="side-panel">
    </div>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    #screen-container {
        @apply flex flex-row;
    }
    .side-panel {
        @apply flex-1 h-auto;
    }
    #game-panel {
        @apply w-[1000px] min-h-screen;
    }
</style>
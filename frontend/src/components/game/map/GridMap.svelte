<script lang="ts">
    import { onMount } from 'svelte';

    // --- Base Settings ---
    let mapSettings = $state.raw({
        mapSize: 1000,    // Total tiles per axis (1000x1000)
        chunkSize: 5      // Tiles per chunk edge
    });

    let currentVillage = $state.raw({
        coordX: 493,
        coordY: 520
    });

    type TileType = "grass" | "water" | "dirt" | "lava";

    interface Tile {
        tile: TileType;
        village?: { player: string; points: number };
    }

    interface Chunk {
        coordX: number;
        coordY: number;
        slots: Tile[];
    }

    let chunks: Chunk[] = $state.raw([]);
    let tilePixelSize = $state.raw(100);
    let chunkPixelSize = $derived(mapSettings.chunkSize * tilePixelSize);
    let chunksPerAxis = $derived(mapSettings.mapSize / mapSettings.chunkSize);

    // --- Tile & Chunk Generation ---
    function generateRandomTile(): Tile {
        const types: TileType[] = ["grass", "water", "dirt", "lava"];
        const tile = types[Math.floor(Math.random() * types.length)];
        let village = undefined;
        if (tile === "grass" && Math.random() < 0.1) {
            village = {
                player: "Player " + Math.floor(Math.random() * 100),
                points: Math.floor(Math.random() * 1000)
            };
        }
        return { tile, village };
    }

    function loadChunk(chunkX: number, chunkY: number) {
        if (chunks.find(c => c.coordX === chunkX && c.coordY === chunkY)) return;
        setTimeout(() => {
            const slots: Tile[] = [];
            const totalSlots = mapSettings.chunkSize * mapSettings.chunkSize;
            for (let i = 0; i < totalSlots; i++) {
                slots.push(generateRandomTile());
            }
            const newChunk: Chunk = { coordX: chunkX, coordY: chunkY, slots };
            chunks = [...chunks, newChunk];
            console.log('Chunk added:', chunkX, chunkY, slots.length);
        }, Math.random() * 300);
    }

    function loadChunks(viewportLeft: number, viewportTop: number, viewportWidth: number, viewportHeight: number) {
        const tileLeft = Math.floor(viewportLeft / tilePixelSize);
        const tileTop = Math.floor(viewportTop / tilePixelSize);
        const tileRight = Math.ceil((viewportLeft + viewportWidth) / tilePixelSize);
        const tileBottom = Math.ceil((viewportTop + viewportHeight) / tilePixelSize);

        const chunkLeft = Math.floor(tileLeft / mapSettings.chunkSize);
        const chunkTop = Math.floor(tileTop / mapSettings.chunkSize);
        const chunkRight = Math.floor(tileRight / mapSettings.chunkSize);
        const chunkBottom = Math.floor(tileBottom / mapSettings.chunkSize);

        for (let cy = chunkTop; cy <= chunkBottom; cy++) {
            for (let cx = chunkLeft; cx <= chunkRight; cx++) {
                if (cx >= 0 && cy >= 0 && cx < chunksPerAxis && cy < chunksPerAxis) {
                    loadChunk(cx, cy);
                }
            }
        }
    }

    let container: HTMLDivElement;
    let scrollTimeout: ReturnType<typeof setTimeout> | null = null;

    function handleScroll() {
        if (scrollTimeout) return;
        scrollTimeout = setTimeout(() => {
            loadChunks(container.scrollLeft, container.scrollTop, container.clientWidth, container.clientHeight);
            scrollTimeout = null;
        }, 100); // throttle
    }

    onMount(() => {
        const centerX = (currentVillage.coordX * tilePixelSize) - container.clientWidth / 2;
        const centerY = (currentVillage.coordY * tilePixelSize) - container.clientHeight / 2;
        container.scrollLeft = centerX;
        container.scrollTop = centerY;
        loadChunks(centerX, centerY, container.clientWidth, container.clientHeight);
    });
</script>

<p class="fixed top-[100px] left-2 z-50 text-white text-sm bg-black bg-opacity-50 p-2 rounded">
    Loaded Chunks: {chunks.length}
</p>

<!-- === Scrollable Map === -->
<div bind:this={container} class="map_window" onscroll={handleScroll}>
    <div
        style="position: relative; width: {mapSettings.mapSize * tilePixelSize}px; height: {mapSettings.mapSize * tilePixelSize}px;"
    >
        {#each chunks as chunk (chunk.coordX + '-' + chunk.coordY)}
            {#if chunk.slots && chunk.slots.length > 0}
                <div
                    class="chunk absolute border border-gray-400 shadow-sm"
                    style="
                        left: {chunk.coordX * chunkPixelSize}px;
                        top: {chunk.coordY * chunkPixelSize}px;
                        width: {chunkPixelSize}px;
                        height: {chunkPixelSize}px;"
                >
                    {#each Array(mapSettings.chunkSize) as _, rowIndex}
                        <div class="row flex">
                            {#each chunk.slots.slice(rowIndex * mapSettings.chunkSize, (rowIndex + 1) * mapSettings.chunkSize) as tile}
                                <div
                                    class="tile border border-gray-300 text-center relative text-[10px] font-semibold flex items-center justify-center hover:ring-2 hover:ring-yellow-300"
                                    style="
                                        width: {tilePixelSize}px;
                                        height: {tilePixelSize}px;
                                        background-color: {
                                            tile.tile === 'grass' ? '#88cc88' :
                                            tile.tile === 'water' ? '#66b2ff' :
                                            tile.tile === 'dirt' ? '#deb887' :
                                            '#ff4500'
                                        };"
                                >
                                    {#if tile.village}
                                        <div class="absolute inset-0 bg-black bg-opacity-40 text-white flex items-center justify-center">
                                            {tile.village.player}
                                        </div>
                                    {/if}
                                </div>
                            {/each}
                        </div>
                    {/each}
                </div>
            {/if}
        {/each}
    </div>
</div>

<style lang="postcss">
    @reference "tailwindcss";

    .map_window {
        @apply overflow-auto h-[800px] w-full bg-gray-500;
    }
</style>
<script lang="ts">
	import { onMount } from "svelte";

	type Village = {
		CoordX: number;
		CoordY: number;
		Player: string;
		Points: number;
	};

	type Chunk = {
		CoordX: number;
		CoordY: number;
		Villages: Village[];
	};

	let settings = $state.raw({
		MapSize: 1000,
		ChunkSize: 5,
		TileSizePx: 100,
		StartPosition: {
			CoordX: 453,
			CoordY: 569,
		},
	});

	let chunks: { [dict_key: string]: Chunk } = $state({});

	let msg: string = $state.raw("");

	function GetChunk(x: number, y: number) {
		const c_x = x - (x % settings.ChunkSize);
		const c_y = y - (y % settings.ChunkSize);
		const coords = `${c_x}|${c_y}`;

		if (chunks[coords]) {
			msg = `Chunk at ${coords} found.`;
			return;
		}

		msg = `Chunk at ${coords} not found – generating it.`;
		const c: Chunk = {
			CoordX: c_x,
			CoordY: c_y,
			Villages: [],
		};

		const vCount = Math.floor(Math.random() * (settings.ChunkSize + 1));
		for (let i = 0; i < vCount; i++) {
			c.Villages.push({
				CoordX: c_x + i,
				CoordY: c_y + i,
				Player: `Player ${coords}-${i}`,
				Points: Math.floor(Math.random() * 10_000),
			});
		}
		chunks[coords] = c;
	}

	let mapWindowEl: HTMLDivElement;

	function generateVisibleChunks() {
		if (!mapWindowEl) return;

		const scrollLeft = mapWindowEl.scrollLeft;
		const scrollTop = mapWindowEl.scrollTop;
		const viewportWidth = mapWindowEl.clientWidth;
		const viewportHeight = mapWindowEl.clientHeight;

		const visibleMinTileX = Math.floor(scrollLeft / settings.TileSizePx);
		const visibleMaxTileX = Math.ceil((scrollLeft + viewportWidth) / settings.TileSizePx);
		const visibleMinTileY = Math.floor(scrollTop / settings.TileSizePx);
		const visibleMaxTileY = Math.ceil((scrollTop + viewportHeight) / settings.TileSizePx);

		for (let tileX = visibleMinTileX; tileX < visibleMaxTileX + settings.ChunkSize; tileX += settings.ChunkSize) {
			for (let tileY = visibleMinTileY; tileY < visibleMaxTileY + settings.ChunkSize; tileY += settings.ChunkSize) {
				GetChunk(tileX, tileY);
			}
		}
	}

	onMount(() => {
		// Calculate the center position in pixels.
		const centerX = settings.StartPosition.CoordX * settings.TileSizePx;
		const centerY = settings.StartPosition.CoordY * settings.TileSizePx;

		if (mapWindowEl) {
			// Scroll so the starting tile is centered.
			mapWindowEl.scrollLeft = centerX - mapWindowEl.clientWidth / 2;
			mapWindowEl.scrollTop = centerY - mapWindowEl.clientHeight / 2;
		}

		generateVisibleChunks();
	});
</script>

<div
	bind:this={mapWindowEl}
	class="map_window"
	onscroll={generateVisibleChunks}
>
	<div
		class="map_grid"
		style="width: {settings.MapSize * settings.TileSizePx}px; height: {settings.MapSize * settings.TileSizePx}px;"
	>
		{#each Object.values(chunks) as chunk (chunk.CoordX + '|' + chunk.CoordY)}
			<div
				class="chunk"
				style="left: {chunk.CoordX * settings.TileSizePx}px; top: {chunk.CoordY * settings.TileSizePx}px;
				       width: {settings.ChunkSize * settings.TileSizePx}px; height: {settings.ChunkSize * settings.TileSizePx}px;"
			>
				{#each chunk.Villages as village (village.CoordX + '|' + village.CoordY)}
					<div
						class="village w-[{settings.TileSizePx}px] h-[{settings.TileSizePx}px]"
						style="left: {(village.CoordX - chunk.CoordX) * settings.TileSizePx}px;
						       top: {(village.CoordY - chunk.CoordY) * settings.TileSizePx}px;"
					>
						<span class="text-xs text-white">{village.Player}</span>
					</div>
				{/each}
			</div>
		{/each}
	</div>
</div>

<p class="mt-4 text-center">{msg}</p>

<style lang="postcss">
    @reference "tailwindcss";

	.map_window {
		@apply relative overflow-auto h-[800px] w-full bg-gray-500;
        -ms-overflow-style: none;
        scrollbar-width: none;
	}
    .map_window::-webkit-scrollbar {
        display: none;
    }

	.map_grid {
		@apply relative;
	}

	.chunk {
		@apply absolute border border-gray-700;
	}

	.village {
		@apply absolute flex items-center justify-center bg-blue-500 rounded-full;
	}


</style>
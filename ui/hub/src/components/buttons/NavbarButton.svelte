<script lang="ts">
	import type { TabName } from '$lib/ts/types/NavbarTypes.svelte';
	import { Button } from 'flowbite-svelte';
	import type { Component } from 'svelte';

	let {
		onclick,
		tabName,
		Icon,
		activeTab
	}: {
		onclick: () => void;
		tabName: TabName;
		Icon: Component;
		activeTab: TabName;
	} = $props();
	let active = $derived(activeTab === tabName);
</script>

<div class="container">
	<div class="panel left-pannel">
		<div class="active-indicator {active ? 'active' : ''}"></div>
		<div class="dummy"></div>
	</div>
	<Button
		class="h-[48px] w-[48px] rounded-2xl bg-transparent focus-within:ring-0 hover:bg-gray-700/90 {active
			? 'bg-gray-700/90'
			: ''}"
		color="dark"
		{onclick}
	>
		<Icon class="scale-140" />
	</Button>
	<div class="pannel right-pannel"></div>
</div>

<style lang="postcss">
	@reference "tailwindcss";
	@theme {
		--shadow-3xl: 0px 0px 70px 15px rgba(0.25, 0.25, 0.25, 0.25);
	}
	.container {
		@apply flex flex-row transition-all duration-300 ease-in-out;
	}
	.pannel {
		@apply h-full w-full;
	}
	.left-pannel {
		@apply flex w-full flex-row;
	}
	.active-indicator {
		@apply mt-[11px] mb-[11px] flex-2 rounded-r-xl;
	}
	.dummy {
		@apply flex-6;
	}
	.active {
		@apply shadow-3xl bg-orange-700 shadow-orange-700;
	}
</style>

<script lang="ts">
	import { Button } from "flowbite-svelte";

    const { onclick, txt } : { onclick: (()=>void) | undefined, txt: string } = $props();
</script>

<Button class="p-0 pl-[5px] pr-[5px] ml-[5px] mr-[5px] rounded-sm bg-gray-500 hover:bg-gray-500 hover:text-green-500 focus-within:ring-0" color="dark" {onclick}>
    {txt}
</Button>

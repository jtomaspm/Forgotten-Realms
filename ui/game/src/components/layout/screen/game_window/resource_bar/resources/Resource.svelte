<script lang="ts">
    import { Button } from 'flowbite-svelte';
	import type { Component } from 'svelte';
    const {
        txt,
        Icon
    } : {
        txt: string
        Icon: Component
    } = $props();
</script>

<Button class="mb-[10px]" color="dark">
    <Icon />
    <p>{txt}</p>
</Button>
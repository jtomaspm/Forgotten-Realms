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

<Button class="h-[35px]" color="dark">
    <div class="mt-[5px] flex items-center">
        <Icon />
        <p>{txt}</p>
    </div>
</Button>
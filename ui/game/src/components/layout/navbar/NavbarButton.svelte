<script lang="ts">
    import { Button } from 'flowbite-svelte';
	import type { Component } from 'svelte';
    let { 
        color,
        href,
        txt,
        Icon
    }:{
        color: "red" | "yellow" | "green" | "purple" | "blue" | "light" | "dark" | "primary" | "none" | "alternative" | undefined,
        href: string,
        txt: string,
        Icon: Component
    } = $props()
</script>

<Button class="p-[10px] focus-within:ring focus-within:ring-black" {color} shadow {href}>
    <Icon />
    <p class={txt.length === 0 ? "" : "ml-[5px]"}>{txt}</p>
</Button>


<style lang="postcss">
    @reference "tailwindcss";
</style>
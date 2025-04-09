<script lang="ts">
	import { TabItem } from "flowbite-svelte";
	import type { Component } from "svelte";

    let {
        name,
        open,
        Icon,
        Content
    }: {
        name: string
        open: boolean
        Icon: Component
        Content: Component
    } = $props();

</script>

<TabItem {open}>
    <div slot="title" class="flex items-center gap-2">
        <Icon />
        {name}
    </div>
        <Content />
</TabItem>
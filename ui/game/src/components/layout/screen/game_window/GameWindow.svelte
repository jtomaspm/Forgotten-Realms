<script lang="ts">
    import ResourceBar from "./resource_bar/ResourceBar.svelte";
    import DisplayContainer from "./display_container/DisplayContainer.svelte";
    let { children } = $props();
    let begginersProtection = true;
</script>

<div id="window-container">
    <ResourceBar />
    {#if begginersProtection}
        <DisplayContainer>
            <p>Begginers Protection ends in x.xx:xx:xx</p>
        </DisplayContainer>
    {/if}
    <DisplayContainer>
        {@render children()}
    </DisplayContainer>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    #window-container {
        @apply w-full mb-[50px] bg-yellow-200;
    }
</style>
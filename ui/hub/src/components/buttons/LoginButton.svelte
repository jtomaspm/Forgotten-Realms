<script lang="ts">
    import { Button } from 'flowbite-svelte';
    function onclick() {
        window.location.href = 'http://localhost:8080/api/github/login';
    }  
</script>

<Button color="dark" {onclick}>
    Sign In
</Button>

<style lang="postcss">
    @reference "tailwindcss";
</style>

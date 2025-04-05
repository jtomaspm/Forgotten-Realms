<script lang="ts">
    import { Button } from 'flowbite-svelte';
    function onclick() {
        window.location.href = 'http://localhost:8080/api/github/login';
    }  
</script>

<div class="profile-btn-container">
    <Button color="dark" {onclick}>
        Login
    </Button>
</div>

<style lang="postcss">
    @reference "tailwindcss";
    .profile-btn-container {
        @apply fixed top-6 right-6 z-50;
    }
</style>

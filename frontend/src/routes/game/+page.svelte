<script lang="ts">
	import { goto } from "$app/navigation";
    $effect.pre(() => {
        goto("/game/village");
    });
</script>
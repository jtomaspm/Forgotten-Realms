<script lang="ts">
	import { goto } from "$app/navigation";
    $effect.pre(() => {
        goto("/village");
    });
</script>
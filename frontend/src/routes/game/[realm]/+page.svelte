<script lang="ts">
	import { goto } from "$app/navigation";
    import { page } from "$app/state";
    $effect.pre(() => {
        goto(`/${page.params.realm}/village`);
    });
</script>
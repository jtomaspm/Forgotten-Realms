<script lang="ts">
	import type { User } from "$lib/ts/types/User.svelte";
	import WorldSelector from "./WorldSelector.svelte";

    let { user }
        : { user: User } 
        = $props();
</script>

<WorldSelector {user} />
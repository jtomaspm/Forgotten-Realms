<script lang="ts">
	import type { PlayableRealm } from "$lib/ts/types/Realm.svelte";
	import FooterLink from "./FooterLink.svelte";
	import RealmSelector from "./RealmSelector.svelte";

    const { realm, realms } : { realm: PlayableRealm, realms: PlayableRealm[] } = $props();
</script>

<div class="footer_container">
    <RealmSelector {realm} {realms} />
    <FooterLink txt={"Logout"} onclick={()=>{window.location.href="/";}} />
</div>


<style lang="postcss">
    @reference "tailwindcss";
    .footer_container {
        @apply w-screen h-[25px] p-[1px] bg-gray-500 fixed bottom-0 left-0 z-50 flex justify-center;
    }
</style>
<script lang="ts">
    import { ButtonGroup } from 'flowbite-svelte';
    import NavBarButton from './NavbarButton.svelte';
    import { 
        GridSolid, 
        HomeSolid,
        GlobeSolid,
        ClipboardListSolid,
        MailBoxSolid,
        AwardSolid,
        UsersGroupSolid,
        CogSolid,
        UserSolid
    } from 'flowbite-svelte-icons';
	import type { PlayableRealm } from '$lib/ts/types/Realm.svelte';
    const color = 'dark';
    const { realm }: { realm: PlayableRealm } = $props();
</script>

<div>
    <ButtonGroup>
        <NavBarButton {color} href={`/game/${realm.name}/player`} txt="" Icon={GridSolid} />
        <NavBarButton {color} href={`/game/${realm.name}/village`} txt="" Icon={HomeSolid} />
    </ButtonGroup>
    <NavBarButton {color} href={`/game/${realm.name}/map`} txt="Map" Icon={GlobeSolid} />
    <NavBarButton {color} href={`/game/${realm.name}/ranking`} txt="Ranking" Icon={AwardSolid} />
    <ButtonGroup>
        <NavBarButton {color} href={`/game/${realm.name}/ally`} txt="" Icon={UsersGroupSolid} />
        <NavBarButton {color} href={`/game/${realm.name}/info/player`} txt="" Icon={UserSolid} />
    </ButtonGroup>
    <ButtonGroup>
        <NavBarButton {color} href={`/game/${realm.name}/report`} txt="3" Icon={ClipboardListSolid} />
        <NavBarButton {color} href={`/game/${realm.name}/mail`} txt="7" Icon={MailBoxSolid} />
    </ButtonGroup>
    <NavBarButton {color} href={`/game/${realm.name}/settings`} txt="Settings" Icon={CogSolid} />
</div>


<style lang="postcss">
    @reference "tailwindcss";
    div {
        @apply w-full h-[45px] p-[5px] bg-gray-500 fixed top-0 left-0 z-50 flex gap-3 justify-center ;
    }
</style>
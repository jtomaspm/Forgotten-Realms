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
    const color = 'dark';
</script>

<div>
    <ButtonGroup>
        <NavBarButton {color} href="/game/player" txt="" Icon={GridSolid} />
        <NavBarButton {color} href="/game/village" txt="" Icon={HomeSolid} />
    </ButtonGroup>
    <NavBarButton {color} href="/game/map" txt="Map" Icon={GlobeSolid} />
    <NavBarButton {color} href="/game/ranking" txt="Ranking" Icon={AwardSolid} />
    <ButtonGroup>
        <NavBarButton {color} href="/game/ally" txt="" Icon={UsersGroupSolid} />
        <NavBarButton {color} href="/game/info/player" txt="" Icon={UserSolid} />
    </ButtonGroup>
    <ButtonGroup>
        <NavBarButton {color} href="/game/report" txt="3" Icon={ClipboardListSolid} />
        <NavBarButton {color} href="/game/mail" txt="7" Icon={MailBoxSolid} />
    </ButtonGroup>
    <NavBarButton {color} href="/settings" txt="Settings" Icon={CogSolid} />
</div>


<style lang="postcss">
    @reference "tailwindcss";
    div {
        @apply w-full h-[45px] p-[5px] bg-gray-500 fixed top-0 left-0 z-50 flex gap-3 justify-center ;
    }
</style>
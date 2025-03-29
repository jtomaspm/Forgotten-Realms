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
        <NavBarButton {color} href="/player" txt="" Icon={GridSolid} />
        <NavBarButton {color} href="/village" txt="" Icon={HomeSolid} />
    </ButtonGroup>
    <NavBarButton {color} href="/map" txt="Map" Icon={GlobeSolid} />
    <NavBarButton {color} href="/ranking" txt="Ranking" Icon={AwardSolid} />
    <ButtonGroup>
        <NavBarButton {color} href="/ally" txt="" Icon={UsersGroupSolid} />
        <NavBarButton {color} href="/info/player" txt="" Icon={UserSolid} />
    </ButtonGroup>
    <ButtonGroup>
        <NavBarButton {color} href="/report" txt="" Icon={ClipboardListSolid} />
        <NavBarButton {color} href="/mail" txt="" Icon={MailBoxSolid} />
    </ButtonGroup>
    <NavBarButton {color} href="/settings" txt="Settings" Icon={CogSolid} />
</div>


<style lang="postcss">
    @reference "tailwindcss";
    div {
        @apply w-full h-[45px] p-[5px] bg-gray-500 fixed top-0 left-0 z-50 flex gap-3 justify-center ;
    }
</style>
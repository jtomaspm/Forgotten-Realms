<script lang="ts">
	import { Button, Card, TabItem, Tabs } from "flowbite-svelte";
	import { AdjustmentsVerticalSolid, GridSolid, UserCircleSolid } from "flowbite-svelte-icons";
	import Faction from "./Faction.svelte";
	import Caldari from "./factions/Caldari.svelte";
	import Varnak from "./factions/Varnak.svelte";
	import Dawnhold from "./factions/Dawnhold.svelte";

</script>

<div class="flex items-center justify-center min-h-screen">
    <Card size="lg" padding="xl">
        <Tabs tabStyle="underline">
            <Faction open={true} name={"Caldari"} Icon={AdjustmentsVerticalSolid} Content={Caldari} />
            <Faction open={false} name={"Varnak"} Icon={AdjustmentsVerticalSolid} Content={Varnak} />
            <Faction open={false} name={"Dawnhold"} Icon={AdjustmentsVerticalSolid} Content={Dawnhold} />
        </Tabs>
        <Button>Select</Button>
    </Card>
</div>


<style lang="postcss">
    @reference "tailwindcss";
</style>


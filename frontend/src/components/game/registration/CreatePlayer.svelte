<script lang="ts">
	import { Card, TabItem, Tabs } from "flowbite-svelte";
	import { AdjustmentsVerticalSolid, GridSolid, UserCircleSolid } from "flowbite-svelte-icons";

</script>

<div class="flex items-center justify-center min-h-screen">
    <Card size="lg" padding="xl">
        <Tabs tabStyle="underline">
        <TabItem open>
            <div slot="title" class="flex items-center gap-2">
                <UserCircleSolid size="md" />
                Profile
            </div>
            <p class="text-sm text-gray-500 dark:text-gray-400">
                <b>Profile:</b>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
            </p>
        </TabItem>
        <TabItem>
            <div slot="title" class="flex items-center gap-2">
                <GridSolid size="md" />
                Dashboard
            </div>
            <p class="text-sm text-gray-500 dark:text-gray-400">
                <b>Dashboard:</b>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
            </p>
        </TabItem>
        <TabItem>
            <div slot="title" class="flex items-center gap-2">
                <AdjustmentsVerticalSolid size="md" />
                Settings
            </div>
            <p class="text-sm text-gray-500 dark:text-gray-400">
                <b>Settings:</b>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
            </p>
        </TabItem>
        </Tabs>
    </Card>
</div>


<style lang="postcss">
    @reference "tailwindcss";
</style>


<script>
    import Layout from "../Shared/Layout.svelte";
    import AppLink from "../Partials/AppLink.svelte";
    import Pager from "../Partials/Pager.svelte";
    import { page } from "@inertiajs/inertia-svelte";
</script>

<svelte:head>
    <title>Fireside</title>
</svelte:head>

<Layout>
    {#each $page.props.links as link (link.ID)}
        <AppLink {link} />
    {:else}
        <div class="px-6 py-4 rounded-md bg-gray-100 dark:bg-gray-600">
            <h1 class="text-2xl mb-6 dark:text-gray-300">
                No results available.
            </h1>
        </div>
    {/each}

    {#if $page.props.links.length !== 0}
        <Pager />
    {/if}
</Layout>

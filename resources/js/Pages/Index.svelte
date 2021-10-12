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
    {/each}
    <Pager />
</Layout>

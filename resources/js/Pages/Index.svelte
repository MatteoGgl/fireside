<script>
    import Layout from "../Shared/Layout.svelte";
    import AppLink from "../Partials/AppLink.svelte";
    import { page } from "@inertiajs/inertia-svelte";
</script>

<svelte:head>
    <title>Links</title>
</svelte:head>

<Layout>
    {#each $page.props.links as link (link.ID)}
        <AppLink {link} />
    {/each}
</Layout>

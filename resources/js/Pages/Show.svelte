<script>
    import Layout from "../Shared/Layout.svelte";
    import AppLink from "../Partials/AppLink.svelte";
    import { page } from "@inertiajs/inertia-svelte";

    $: link = $page.props.link;
</script>

<svelte:head>
    <title>{link.Title}</title>
</svelte:head>

<Layout>
    <AppLink {link} expanded={true} />
</Layout>

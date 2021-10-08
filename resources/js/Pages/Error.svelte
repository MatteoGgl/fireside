<script>
    import Layout from "./Layout.svelte";
    import { page } from "@inertiajs/inertia-svelte";
</script>

<svelte:head>
    <title>Error</title>
</svelte:head>

<Layout>
    <main>
        <div class="columns">
            <div class="column has-background-white py-3">
                <p>Error: {$page.props.error}</p>
            </div>
        </div>
    </main>
</Layout>

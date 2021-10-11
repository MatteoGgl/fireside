<script>
    import Layout from "../Shared/Layout.svelte";
    import { page } from "@inertiajs/inertia-svelte";
</script>

<svelte:head>
    <title>Error</title>
</svelte:head>

<Layout>
    <div class="px-6 py-4 bg-gray-100 dark:bg-gray-600">
        <h1 class="text-4xl mb-6 dark:text-gray-300">
            Error: {$page.props.error}
        </h1>
    </div>
</Layout>

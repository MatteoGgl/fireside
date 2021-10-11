<script>
    import { Inertia } from "@inertiajs/inertia";
    import { inertia, page } from "@inertiajs/inertia-svelte";
    import { route } from "../stores.js";
    import DarkMode from "./DarkMode.svelte";

    function logout() {
        Inertia.post(route("auth.logout"));
    }
</script>

<nav
    class="max-w-7xl mx-auto py-6 flex items-center"
    role="navigation"
    aria-label="main navigation"
>
    <a use:inertia class="text-2xl dark:text-gray-300" href={route("index")}>
        <img src="/static/images/logo.png" alt="Fireside logo" class="h-12 w-12">
    </a>
    <div class="ml-auto flex items-center space-x-4">
        {#if $page.props.isAuthenticated}
            <a
                href={route("links.new")}
                class="text-sm hover:text-blue-500 hover:underline dark:text-gray-300 dark:hover:text-indigo-400"
                use:inertia
            >
                New post
            </a>
            <form on:submit|preventDefault={logout}>
                <button
                    type="submit"
                    class="text-sm hover:text-blue-500 hover:underline dark:text-gray-300 dark:hover:text-indigo-400"
                >
                    Logout
                </button>
            </form>
        {:else}
            <a
                href={route("auth.login")}
                class="text-sm hover:text-blue-500 hover:underline dark:text-gray-300 dark:hover:text-indigo-400"
                use:inertia
            >
                Log in
            </a>
        {/if}
        <DarkMode />
    </div>
</nav>

<script>
    import { Inertia } from "@inertiajs/inertia";
    import { inertia, page } from "@inertiajs/inertia-svelte";
    import { route } from "../stores.js"
    import DarkMode from "./DarkMode.svelte"

    function logout() {
        Inertia.post(route("auth.logout"));
    }
</script>

<nav
    class="navbar has-background-light px-6 py-3"
    role="navigation"
    aria-label="main navigation"
>
    <div class="navbar-brand">
        <div class="navbar-item">
            <a use:inertia class="is-size-4" href="{route('index')}">Linki</a>
        </div>
    </div>
    <div class="navbar-item ml-auto">
        {#if $page.props.isAuthenticated}
            <a href="{route('links.new')}" class="mr-3 button is-ghost" use:inertia>Post</a>
            <form action="POST" on:submit|preventDefault={logout}>
                <button type="submit" class="button is-ghost">Logout</button>
            </form>
        {:else}
            <a href="{route('auth.login')}" class="button is-ghost" use:inertia>Log in</a>
        {/if}
        <DarkMode />
    </div>
</nav>

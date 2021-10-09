<script>
    import { inertia } from "@inertiajs/inertia-svelte";
</script>

<nav
    class="navbar has-background-light px-6 py-3"
    role="navigation"
    aria-label="main navigation"
>
    <div class="navbar-brand">
        <div class="navbar-item">
            <a use:inertia class="is-size-4" href="/">Linki</a>
        </div>
    </div>
    <div class="navbar-item ml-auto">
        <a href="/login" use:inertia>Log in</a>
    </div>
</nav>

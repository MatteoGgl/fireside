<script>
    import Navbar from "./Navbar.svelte";
</script>

<Navbar />
<div class="container max-w-7xl mx-auto mt-3">
    <main class="grid gap-4 grid-cols-3">
        <div class="col-span-2">
            <slot />
        </div>
        <div class="col-span-1">Sidebar</div>
    </main>
</div>

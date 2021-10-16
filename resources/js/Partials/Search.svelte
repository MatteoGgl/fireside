<script>
    import { Inertia } from "@inertiajs/inertia";
    import { route } from "../stores.js";

    const qs = new URLSearchParams(window.location.search);

    let search = qs.get("title") ?? "";
    let loading = false;

    function submit() {
        Inertia.visit(
            route("index", {
                title: search,
            }),
            {
                onStart: () => (loading = true),
                onFinish: () => (loading = false),
            }
        );
    }

    function resetSearch() {
        search = "";
        Inertia.visit(route("index"));
    }
</script>

<div class="relative flex items-center w-full lg:w-64 h-full group">
    <svg
        class="absolute left-0 z-20 w-4 h-4 ml-4 text-gray-400 pointer-events-none fill-current"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 20 20"
    >
        <path
            d="M12.9 14.32a8 8 0 1 1 1.41-1.41l5.35 5.33-1.42 1.42-5.33-5.34zM8 14A6 6 0 1 0 8 2a6 6 0 0 0 0 12z"
        />
    </svg>
    <input
        bind:value={search}
        on:keydown={(e) => {
            if (e.key === "Enter") submit();
        }}
        type="search"
        class="block w-full py-1.5 px-10 pr-4 leading-normal rounded-2xl focus:border-transparent focus:outline-none focus:ring-2 focus:ring-blue-500 ring-opacity-90 bg-gray-100 dark:bg-gray-800 dark:text-gray-400 aa-input"
        placeholder="Search"
    />
    <svg
        xmlns="http://www.w3.org/2000/svg"
        class="{loading || search === ""
            ? 'hidden'
            : ''} absolute right-0 z-20 w-4 h-4 mr-4 text-gray-400 fill-current cursor-pointer"
        on:click={resetSearch}
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
    >
        <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M6 18L18 6M6 6l12 12"
        />
    </svg>
    <svg
        class="{loading
            ? ''
            : 'hidden'} absolute right-0 z-20 w-4 h-4 mr-4 text-gray-400 pointer-events-none fill-current animate-spin"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
    >
        <path
            class="opacity-75"
            fill="currentColor"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
        />
    </svg>
</div>

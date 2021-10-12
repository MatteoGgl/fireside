<script>
    import { Inertia } from "@inertiajs/inertia"
    import { page } from "@inertiajs/inertia-svelte";
    import { route } from "../stores.js";

    $: current_page = $page.props.pagedata.current_page;
    $: first_page = $page.props.pagedata.first_page;
    $: last_page = $page.props.pagedata.last_page;

    function nextPage() {
        let next_page = parseInt(current_page) + 1;
        if (next_page > last_page) {
            next_page = last_page;
        }
        Inertia.visit(route("index") + "?page=" + next_page, {
            only: ["links", "pagedata"]
        })
    }

    function previousPage() {
        let prev_page = parseInt(current_page) - 1;
        if (prev_page < first_page) {
            prev_page = first_page;
        }
        Inertia.visit(route("index") + "?page=" + prev_page, {
            only: ["links", "pagedata"]
        })
    }
</script>

<div class="bg-white dark:bg-gray-700 py-5 flex xs:flex-row items-center justify-between">
    <button
        on:click="{previousPage}"
        class="p-4 border dark:border-transparent text-base rounded-md dark:text-gray-300 bg-gray-100 dark:bg-gray-600"
        disabled={current_page === first_page}
    >
        <svg
            width="9"
            fill="currentColor"
            height="8"
            class=""
            viewBox="0 0 1792 1792"
            xmlns="http://www.w3.org/2000/svg"
        >
            <path
                d="M1427 301l-531 531 531 531q19 19 19 45t-19 45l-166 166q-19 19-45 19t-45-19l-742-742q-19-19-19-45t19-45l742-742q19-19 45-19t45 19l166 166q19 19 19 45t-19 45z"
            />
        </svg>
    </button>
    <button
        on:click="{nextPage}"
        type="button"
        class="p-4 border dark:border-transparent text-base rounded-md dark:text-gray-300 bg-gray-100 dark:bg-gray-600"
        disabled={current_page === last_page}
    >
        <svg
            width="9"
            fill="currentColor"
            height="8"
            class=""
            viewBox="0 0 1792 1792"
            xmlns="http://www.w3.org/2000/svg"
        >
            <path
                d="M1363 877l-742 742q-19 19-45 19t-45-19l-166-166q-19-19-19-45t19-45l531-531-531-531q-19-19-19-45t19-45l166-166q19-19 45-19t45 19l742 742q19 19 19 45t-19 45z"
            />
        </svg>
    </button>
</div>

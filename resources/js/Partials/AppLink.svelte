<script>
    import { Inertia } from "@inertiajs/inertia";
    import { inertia, Link, page } from "@inertiajs/inertia-svelte";
    import { route } from "../stores.js";

    export let link;
    export let expanded;
    let deleteConfirm = false;

    function deleteLink() {
        Inertia.delete(route("links.destroy", { id: link.ID }), {
            preserveScroll: true,
        });
    }
</script>

<article
    class="flex items-center px-6 py-4 rounded-md mb-6 relative bg-gray-100 dark:bg-gray-600"
>
    <div class="flex flex-col items-center mr-6 dark:text-gray-300">
        <span>{link.Likes}</span>
        <span>votes</span>
    </div>
    <div class="flex flex-col w-full">
        <a
            class="text-lg font-bold dark:text-gray-300"
            href={route("links.show", { id: link.ID })}
            use:inertia
        >
            {link.Title}
        </a>
        {#if link.Type === "link"}
            <a
                class="mt-1"
                href={link.URL.String}
                target="_blank"
                rel="noopener nofollow noreferrer"
            >
                <span
                    class="text-xs text-gray-500 dark:text-gray-400 flex items-center"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="h-4 w-4 mr-1"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
                        />
                    </svg>
                    {link.URL.String}
                </span>
            </a>
        {/if}
        {#if link.Type === "text"}
            <p
                class="mt-1 {!expanded
                    ? 'truncate max-w-prose text-gray-500 dark:text-gray-400'
                    : 'dark:text-gray-300'} cursor-pointer"
                on:click={() => (expanded = !expanded)}
            >
                {#if expanded}
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="h-4 w-4 inline"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M19 9l-7 7-7-7"
                        />
                    </svg>
                {:else}
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="h-4 w-4 inline"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M9 5l7 7-7 7"
                        />
                    </svg>
                {/if}

                {link.Content.String}
            </p>
        {/if}
        <div class="flex items-center space-x-3 mt-4">
            <span class="text-xs text-gray-500 dark:text-gray-300">
                {link.HumanCreatedAt}
            </span>
            {#each link.Tags as tag}
                <Link
                    href={route("index", {
                        tags: tag,
                    })}
                >
                    <span
                        class="px-2 py-1 text-xs text-white rounded bg-blue-500 dark:bg-indigo-500 dark:text-gray-300"
                    >
                        {tag}
                    </span>
                </Link>
            {/each}
        </div>
        {#if $page.props.isAuthenticated}
            <div
                class="flex items-center justify-end space-x-3 text-xs text-gray-400"
            >
                {#if deleteConfirm}
                    <div class="flex items-center space-x-3">
                        <span class="dark:text-gray-300">are you sure?</span>
                        <form on:submit|preventDefault={deleteLink}>
                            <button
                                type="submit"
                                class="hover:underline hover:text-blue-500 dark:text-gray-300  dark:hover:text-indigo-400"
                            >
                                yes
                            </button>
                        </form>
                        <button
                            on:click={() => (deleteConfirm = false)}
                            class="hover:underline hover:text-blue-500 dark:text-gray-300  dark:hover:text-indigo-400"
                        >
                            no
                        </button>
                    </div>
                {:else}
                    <button
                        on:click={() => (deleteConfirm = true)}
                        class="hover:underline hover:text-blue-500 dark:text-gray-300  dark:hover:text-indigo-400"
                    >
                        delete
                    </button>
                {/if}
            </div>
        {/if}
    </div>
</article>

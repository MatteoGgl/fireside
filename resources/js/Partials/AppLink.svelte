<script>
    import { inertia, Link } from "@inertiajs/inertia-svelte";
    import { route } from "../stores.js";

    export let link;
    export let expanded;
</script>

<article
    class="flex items-center px-6 py-4 rounded-md mb-6 relative bg-gray-100 dark:bg-gray-600"
>
    <div class="flex flex-col items-center mr-6 dark:text-gray-300">
        <span>{link.Likes}</span>
        <span>votes</span>
    </div>
    <div class="flex flex-col">
        <a
            class="text-lg font-bold dark:text-gray-300"
            href={route("links.show", { id: link.ID })}
            use:inertia
        >
            {link.Title}
        </a>
        {#if link.Type === "link"}
            <a
                href={link.URL.String}
                target="_blank"
                rel="noopener nofollow noreferrer"
            >
                <span class="text-xs text-gray-500 dark:text-gray-400">
                    {link.URL.String}
                </span>
            </a>
        {/if}
        {#if link.Type === "text"}
            <p
                class="mt-3 {!expanded
                    ? 'truncate max-w-prose text-gray-500 dark:text-gray-400'
                    : 'dark:text-gray-300'} cursor-pointer"
                on:click={() => (expanded = !expanded)}
            >
                {link.Content.String}
            </p>
        {/if}
        <div class="flex items-center space-x-3 mt-4">
            {#each link.Tags as tag}
                <Link
                    href={route("links.tag", {
                        tag: tag,
                    })}
                >
                    <span
                        class="px-2 py-1 text-xs text-white rounded bg-blue-500 dark:bg-indigo-500 dark:text-gray-300"
                    >
                        {tag}
                    </span>
                </Link>
            {/each}
            <span class="text-xs text-gray-500">
                {link.CreatedAt}
            </span>
        </div>
    </div>
</article>

<script>
    import Layout from "./Layout.svelte";
    import { page, Link } from "@inertiajs/inertia-svelte";
    import { route } from "../stores.js"

    $: link = $page.props.link;
</script>

<svelte:head>
    <title>{link.Title}</title>
</svelte:head>

<Layout>
    <main>
        <div class="columns">
            <div class="column has-background-white py-3 is-four-fifths">
                <h2 class="title">{link.Title}</h2>
                <article class="media">
                    <div
                        class="media-left is-flex is-flex-direction-column is-align-items-center"
                    >
                        <span>{link.Likes}</span>
                        <span>votes</span>
                    </div>
                    <div class="media-content">
                        <div class="content">
                            {#if link.Type === "text"}
                                <p>{link.Content.String}</p>
                            {:else}
                                <a
                                    href={link.URL.String}
                                    target="_blank"
                                    rel="noopener noreferrer nofollow"
                                >
                                    <strong>{link.URL.String}</strong>
                                </a>
                            {/if}

                            <nav class="level my-4">
                                <div class="level-left">
                                    {#if link.Tags.length}
                                        {#each link.Tags as tag}
                                            <div class="level-item">
                                                <Link href={route("links.tag", {tag: tag})}>
                                                    <span class="tag"
                                                        >{tag}</span
                                                    >
                                                </Link>
                                            </div>
                                        {/each}
                                    {/if}
                                </div>
                            </nav>

                            <nav class="level">
                                <div class="level-left">
                                    <div class="level-item">
                                        {link.CreatedAt}
                                    </div>
                                </div>
                            </nav>
                        </div>
                    </div>
                </article>
            </div>
            <div class="column">Sidebar</div>
        </div>
    </main>
</Layout>

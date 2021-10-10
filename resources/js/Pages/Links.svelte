<script>
    import Layout from "./Layout.svelte";
    import { inertia, page, Link } from "@inertiajs/inertia-svelte";
    import { route } from "../stores.js"
</script>

<svelte:head>
    <title>Links</title>
</svelte:head>

<Layout>
    <main>
        <div class="columns">
            <div class="column has-background-white py-3 is-four-fifths">
                {#each $page.props.links as link (link.ID)}
                    <article class="media">
                        <div
                            class="media-left is-flex is-flex-direction-column is-align-items-center"
                        >
                            <span>{link.Likes}</span>
                            <span>votes</span>
                        </div>
                        <div class="media-content">
                            <div class="content">
                                <a href="{route('links.show', {'id': link.ID})}" use:inertia>
                                    <strong>{link.Title}</strong>
                                </a>
                                {#if link.Type === "link"}
                                    <a
                                        href={link.URL.String}
                                        target="_blank"
                                        rel="noopener nofollow noreferrer"
                                    >
                                        <span
                                            class="is-size-7 has-text-grey tag is-link is-light"
                                        >
                                            {link.URL.String}
                                        </span>
                                    </a>
                                {/if}
                                <nav class="level mt-3">
                                    <div class="level-left">
                                        {#each link.Tags as tag}
                                            <div class="level-item">
                                                <Link href="/tag/{tag}">
                                                    <span class="tag"
                                                        >{tag}</span
                                                    >
                                                </Link>
                                            </div>
                                        {/each}
                                        <div class="level-item">
                                            {link.CreatedAt}
                                        </div>
                                    </div>
                                </nav>
                            </div>
                        </div>
                    </article>
                {/each}
            </div>
            <div class="column">Sidebar</div>
        </div>
    </main>
</Layout>

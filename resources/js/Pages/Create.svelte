<script>
    import Layout from "./Layout.svelte";
    import { useForm } from "@inertiajs/inertia-svelte";
    import { route } from "../stores.js";

    let tab = "link";

    let form = useForm({
        title: null,
        type: null,
        url: null,
        content: null,
        tags: null,
    });

    function otherTab() {
        tab = tab === "link" ? "text" : "link";
        $form.clearErrors();
    }

    function submit() {
        $form
            .transform((data) => ({
                ...data,
                type: tab,
            }))
            .post(route("links.store"));
    }
</script>

<svelte:head>
    <title>Create</title>
</svelte:head>

<Layout>
    <main>
        <div class="columns is-centered">
            <div class="column is-6 has-background-white py-3">
                <h1 class="title">Create new link</h1>

                <div class="tabs is-boxed">
                    <!-- svelte-ignore a11y-missing-attribute -->
                    <ul>
                        <li class={tab === "link" ? "is-active" : ""}>
                            <a on:click|preventDefault={otherTab}> Link </a>
                        </li>
                        <li class={tab === "text" ? "is-active" : ""}>
                            <a on:click|preventDefault={otherTab}> Text </a>
                        </li>
                    </ul>
                </div>

                <form on:submit|preventDefault={submit} action="POST">
                    <div class="field">
                        <label class="label" for="title">Title *</label>
                        <div class="control">
                            <input
                                id="title"
                                name="title"
                                class="input"
                                type="text"
                                placeholder="My awesome new link"
                                bind:value={$form.title}
                                error="$form.errors.title"
                            />
                        </div>
                        {#if $form.errors.title}
                            <p class="help is-danger">
                                {$form.errors.title}
                            </p>
                        {/if}
                    </div>

                    {#if tab === "link"}
                        <div class="field">
                            <label class="label" for="title">URL *</label>
                            <div class="control">
                                <input
                                    id="url"
                                    name="url"
                                    class="input"
                                    type="url"
                                    placeholder="https://..."
                                    bind:value={$form.url}
                                    error="$form.errors.url"
                                />
                            </div>
                            {#if $form.errors.url}
                                <p class="help is-danger">
                                    {$form.errors.url}
                                </p>
                            {/if}
                        </div>
                    {:else if tab === "text"}
                        <div class="field">
                            <label class="label" for="content">Content *</label>
                            <div class="control">
                                <textarea
                                    class="textarea"
                                    name="content"
                                    id="content"
                                    rows="5"
                                    bind:value={$form.content}
                                />
                            </div>
                            {#if $form.errors.content}
                                <p class="help is-danger">
                                    {$form.errors.content}
                                </p>
                            {/if}
                        </div>
                    {/if}

                    <div class="field">
                        <label class="label" for="tags">
                            Tags (separated by comma)
                        </label>
                        <div class="control">
                            <input
                                id="tags"
                                name="tags"
                                class="input"
                                type="text"
                                placeholder="first tag, second tag, third tag"
                                bind:value={$form.tags}
                                error="$form.errors.tags"
                            />
                        </div>
                        {#if $form.errors.tags}
                            <p class="help is-danger">
                                {$form.errors.tags}
                            </p>
                        {/if}
                    </div>

                    <div class="field is-grouped">
                        <div class="control">
                            <button
                                type="submit"
                                disabled={$form.processing}
                                class="button is-link"
                            >
                                Submit
                            </button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </main>
</Layout>

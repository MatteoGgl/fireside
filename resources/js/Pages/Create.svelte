<script>
    import Layout from "../Shared/Layout.svelte";
    import Input from "../Partials/Input.svelte";
    import Textarea from "../Partials/Textarea.svelte";
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
    <div class="px-6 py-4 rounded-md bg-gray-100 dark:bg-gray-600">
        <h1 class="text-4xl mb-6 dark:text-gray-300">Create new link</h1>

        <div class="flex space-x-4 border-b-2 dark:border-indigo-300 items-end mb-4">
            <button
                class="-mb-px pb-2 px-3 {tab === 'link'
                    ? 'font-bold text-blue-500 dark:text-indigo-400 border-b-2 border-blue-500 dark:border-indigo-500'
                    : 'font-normal text-black hover:text-blue-500 dark:text-gray-300'}"
                on:click={otherTab}>Link</button
            >
            <button
                class="-mb-px pb-2 px-3 {tab === 'text'
                    ? 'font-bold text-blue-500 dark:text-indigo-400 border-b-2 border-blue-500 dark:border-indigo-500'
                    : 'font-normal text-black hover:text-blue-500 dark:text-gray-300'}"
                on:click={otherTab}>Text</button
            >
        </div>
        <form on:submit|preventDefault={submit}>
            <div class="flex flex-col max-w-lg mt-6 space-y-5">
                <Input
                    type="text"
                    form={$form}
                    name="title"
                    label="Title"
                    placeholder="My awesome new link"
                    required={true}
                />
                {#if tab === "link"}
                    <Input
                        type="url"
                        form={$form}
                        name="url"
                        label="URL"
                        placeholder="https://..."
                        required={true}
                    />
                {:else}
                    <Textarea
                        form={$form}
                        name="content"
                        label="Contents"
                        required={true}
                    />
                {/if}
                <Input
                    type="text"
                    form={$form}
                    name="tags"
                    label="Tags (separated by comma)"
                    placeholder="first tag, second tag, third tag..."
                />
                <button
                    type="submit"
                    disabled={$form.processing}
                    class="py-2 px-4 text-center rounded-md text-white bg-blue-500 dark:bg-indigo-500 hover:bg-blue-600 dark:hover:bg-indigo-600 mr-auto"
                >
                    Submit
                </button>
            </div>
        </form>
    </div>
</Layout>

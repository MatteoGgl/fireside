<script>
    import Layout from "../Shared/Layout.svelte";
    import Input from "../Partials/Input.svelte";
    import { useForm } from "@inertiajs/inertia-svelte";
    import { route } from "../stores.js";

    let form = useForm({
        email: null,
        password: null,
    });

    function submit() {
        $form.post(route("auth.loginPost"));
    }
</script>

<svelte:head>
    <title>Log in</title>
</svelte:head>

<Layout>
    <div class="px-6 py-4 rounded-md bg-gray-100 dark:bg-gray-600">
        <h1 class="text-4xl mb-6 dark:text-gray-300">Login</h1>

        <form on:submit|preventDefault={submit}>
            <div class="flex flex-col max-w-lg mt-6 space-y-5">
                <Input
                    type="email"
                    form={$form}
                    name="email"
                    label="Email"
                    placeholder="test@example.com"
                    required={true}
                />

                {#if $form.errors.generic}
                    <p class="text-sm text-red-500 dark:text-red-400">
                        {$form.errors.generic}
                    </p>
                {/if}

                <Input
                    type="password"
                    form={$form}
                    name="password"
                    label="Password"
                    placeholder="Enter your secret password"
                    required={true}
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

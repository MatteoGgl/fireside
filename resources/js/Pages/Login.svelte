<script>
    import Layout from "./Layout.svelte";
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
    <main>
        <div class="columns is-centered">
            <div class="column is-6 has-background-white py-3">
                <h1 class="title">Login</h1>

                <form on:submit|preventDefault={submit} action="POST">
                    <div class="field">
                        <label class="label" for="email">Email</label>
                        <div class="control">
                            <input
                                id="email"
                                name="email"
                                class="input"
                                type="email"
                                placeholder="Your email"
                                bind:value={$form.email}
                                error="$form.errors.email"
                            />
                        </div>
                        {#if $form.errors.email}
                            <p class="help is-danger">
                                {$form.errors.email}
                            </p>
                        {/if}
                        {#if $form.errors.generic}
                            <p class="help is-danger">
                                {$form.errors.generic}
                            </p>
                        {/if}
                    </div>
                    <div class="field">
                        <label class="label" for="password">Password</label>
                        <div class="control">
                            <input
                                id="password"
                                class="input"
                                type="password"
                                placeholder="Your password"
                                bind:value={$form.password}
                                error="$form.errors.password"
                            />
                        </div>
                        {#if $form.errors.password}
                            <p class="help is-danger">
                                {$form.errors.password}
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

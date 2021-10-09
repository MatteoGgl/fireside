<script>
    import Layout from "./Layout.svelte";
    import { page } from "@inertiajs/inertia-svelte";
</script>

<svelte:head>
    <title>Log in</title>
</svelte:head>

<Layout>
    <main>
        <div class="columns is-centered">
            <div class="column is-6 has-background-white py-3">
                <h1 class="title">Login</h1>

                <form action="POST">
                    <div class="field">
                        <label class="label" for="email">Email</label>
                        <div class="control">
                            <input
                                id="email"
                                name="email"
                                class="input"
                                type="email"
                                placeholder="Your email"
                            />
                        </div>
                    </div>
                    <div class="field">
                        <label class="label" for="password">Password</label>
                        <div class="control">
                            <input
                                id="password"
                                class="input"
                                type="password"
                                placeholder="Your password"
                            />
                        </div>
                    </div>
                    <div class="field is-grouped">
                        <div class="control">
                            <button type="submit" class="button is-link">
                                Submit
                            </button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </main>
</Layout>

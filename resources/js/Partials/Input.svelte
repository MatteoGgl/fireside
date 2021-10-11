<script>
    import { onMount } from "svelte";

    export let type = "text";
    export let form;
    export let name;
    export let label;
    export let placeholder;
    export let required = false;

    let inputElement;
    onMount(() => {
        inputElement.type = type;
    });
</script>

<div class="relative">
    <label for={name} class="dark:text-gray-300">
        {label}
        {#if required}
            <span class="text-red-600 dark:text-red-400">*</span>
        {/if}
    </label>
    <input
        bind:this={inputElement}
        id={name}
        {name}
        {placeholder}
        bind:value={form[name]}
        class="mt-3 rounded-lg border-transparent flex-1 appearance-none border border-gray-300 dark:border-transparent w-full py-2 px-4 bg-white dark:bg-gray-500 text-gray-700 dark:text-gray-300 placeholder-gray-400 shadow-sm text-base focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-indigo-500 focus:border-transparent"
    />
    {#if form.errors[name]}
        <p class="text-sm text-red-500 dark:text-red-400">
            {form.errors[name]}
        </p>
    {/if}
</div>

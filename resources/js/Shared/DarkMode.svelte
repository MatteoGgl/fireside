<script>
    import { onMount } from "svelte";

    onMount(() => {
        applyTheme();
        window
            .matchMedia(DARK_PREFERENCE)
            .addEventListener("change", applyTheme);

        return () => window.removeEventListener("change", applyTheme);
    });

    const STORAGE_KEY = "theme";
    const DARK_PREFERENCE = "(prefers-color-scheme: dark)";
    const THEMES = {
        DARK: "dark",
        LIGHT: "light",
    };
    let currentTheme = undefined;

    const prefersDarkTheme = () => window.matchMedia(DARK_PREFERENCE).matches;
    const toggleTheme = () => {
        const stored = localStorage.getItem(STORAGE_KEY);

        if (stored) {
            localStorage.removeItem(STORAGE_KEY);
        } else {
            localStorage.setItem(
                STORAGE_KEY,
                prefersDarkTheme() ? THEMES.LIGHT : THEMES.DARK
            );
        }

        applyTheme();
    };

    const applyTheme = () => {
        const preferredTheme = prefersDarkTheme() ? THEMES.DARK : THEMES.LIGHT;
        currentTheme = localStorage.getItem(STORAGE_KEY) ?? preferredTheme;

        if (currentTheme === THEMES.DARK) {
            document.documentElement.classList.remove(THEMES.LIGHT);
            document.documentElement.classList.add(THEMES.DARK);
        } else {
            document.documentElement.classList.remove(THEMES.DARK);
            document.documentElement.classList.add(THEMES.LIGHT);
        }
    };
</script>

<div class="relative inline-block w-10 align-middle select-none">
    <input
        id="toggle"
        name="toggle"
        type="checkbox"
        on:click={toggleTheme}
        checked={currentTheme === THEMES.DARK}
        class="checked:bg-gray-500 outline-none focus:outline-none right-4 checked:right-0 duration-200 ease-in absolute block w-6 h-6 rounded-full bg-white border-4 appearance-none cursor-pointer"
    />
    <label
        for="toggle"
        class="block overflow-hidden h-6 rounded-full bg-gray-300 cursor-pointer"
    />
</div>

<script>
    import { onMount } from "svelte";

    onMount(() => {
        applyTheme();
        window
            .matchMedia(DARK_PREFERENCE)
            .addEventListener("change", applyTheme);

        return () => window.removeEventListener("change");
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
            document.body.classList.remove(THEMES.LIGHT);
            document.body.classList.add(THEMES.DARK);
        } else {
            document.body.classList.remove(THEMES.DARK);
            document.body.classList.add(THEMES.LIGHT);
        }
    };
</script>

<button class="button" on:click={toggleTheme}
    >{currentTheme === THEMES.DARK ? "Light" : "Dark"}</button
>

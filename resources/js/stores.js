import { derived, get } from 'svelte/store'
import { page } from "@inertiajs/inertia-svelte"

const appRoutes = derived(page, $page => JSON.parse($page.props.routes));

export const route = (name, params = {}) => {
    var route = get(appRoutes).find((r) => {
        return r.name === name;
    });
    if (typeof route === "undefined") {
        console.error("cannot find route:", name);
        return "";
    }

    var path = route.path;
    Object.entries(params).forEach(([key, val]) => {
        path = path.replace(":" + key, val)
    });
    return path;
};

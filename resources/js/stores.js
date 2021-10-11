/*
* stores.js
* Copyright (C) <2021>  <Matteo Guglielmetti>
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU Affero General Public License as published
* by the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU Affero General Public License for more details.
*
* You should have received a copy of the GNU Affero General Public License
* along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

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

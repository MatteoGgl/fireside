const mix = require('laravel-mix');
const path = require('path');

require('laravel-mix-svelte');

mix.js('resources/js/app.js', 'ui/static/js')
    .sass('resources/css/app.scss', 'ui/static/css')
    .svelte({
        dev: !mix.inProduction()
    })
    .disableNotifications();
const mix = require('laravel-mix');
const path = require('path');

require('laravel-mix-svelte');

mix.js('resources/js/app.js', 'ui/static/js')
    .css('resources/css/app.css', 'ui/static/css')
    .svelte({
        dev: !mix.inProduction()
    })
    .disableNotifications();
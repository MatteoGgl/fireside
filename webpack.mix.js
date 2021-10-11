const mix = require('laravel-mix');

require('laravel-mix-svelte');

mix.js('resources/js/app.js', 'ui/static/js')
    .postCss('resources/css/app.css', 'ui/static/css', [
        require('tailwindcss'),
        require("autoprefixer")
    ])
    .svelte({
        dev: !mix.inProduction()
    })
    .disableNotifications();
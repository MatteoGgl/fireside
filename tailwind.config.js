module.exports = {
  purge: {
    //enabled: true,
    content: [
      './resources/js/*/*.svelte',
      './ui/app.tmpl'
    ]
  },
  darkMode: "class",
  theme: {
    extend: {},
  },
  variants: {
    extend: {
      backgroundColor: ["checked"],
      borderColor: ["checked"],
      inset: ["checked"],
      zIndex: ["hover", "active"],
    },
  },
  plugins: [],
  future: {
    purgeLayersByDefault: true,
  }
}

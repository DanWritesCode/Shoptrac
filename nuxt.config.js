export default {
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'stonksup',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  /* proxy the golang API server while in development */
  proxy: {
    '/api/': { target: 'http://localhost:8000', pathRewrite: {'^/api/': ''} }
  },

  /*
   ** Customize the progress-bar color
   */
  loading: { color: '#EF423C' },
  /*
   ** Global CSS
   */
  css: [
    // SCSS file in the project
    '@/assets/sass/app.scss',
    '@/assets/sass/bootstrap.scss',
    '@fortawesome/fontawesome-svg-core/styles.css'
  ],
  /*
   ** Plugins to load before mounting the App
   */
  plugins: [
    '~/plugins/vee-validate',
    '~/plugins/fontawesome',
    '~/plugins/v-transitions',
    '~/plugins/vue-sanitize',

  ],
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [],
  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    // Doc: https://github.com/nuxt-community/dotenv-module
    '@nuxtjs/dotenv',
  ],
  /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
  axios: {
    proxy: true
  },
  /*
   ** Build configuration
   */
  build: {
    parallel: true,
    transpile: ['vee-validate/dist/rules'],
    splitChunks: {
      layouts: true
    },
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {}
  },
  minimize: true,
  minimizer: [
    // terser-webpack-plugin
    // optimize-css-assets-webpack-plugin
  ],
  splitChunks: {
    chunks: 'all',
    automaticNameDelimiter: '.',
    name: undefined,
    cacheGroups: {}
  }
}

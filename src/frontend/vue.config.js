const path = require('path')
const PrerenderSpaPlugin = require('prerender-spa-plugin')
const Renderer = PrerenderSpaPlugin.PuppeteerRenderer

module.exports = {
  publicPath: process.env.BASE_URL,

  // So we don't get an 'Invalid Host' error message when we
  // try to run this in production mode. Notes here: https://bit.ly/3lI7sig
  // don't set the scheme (duh) because it's a host, not a fqdn
  devServer: {
    public: 'skillcalc5e.cc'
  },

  // only done for the title right now
  pages: {
    index: {
      entry: 'src/main.ts',
      title: '5e Skill Challenge Calculator'
    }
  },

  // there might be a better way to handle this in the future, (Vue 3 launch)
  // but this sets up pre-rendering the site for SEO
  chainWebpack: config => {
    config
        .plugin("prerender-spa-plugin")
        .use(PrerenderSpaPlugin)
        .init(Plugin => new Plugin({
            staticDir: path.join(__dirname, 'dist'),
            routes: ['/'],
            postProcess(renderedRoute) {
              renderedRoute.html = renderedRoute.html.replace(`id="app"`, `id="app" data-server-rendered="true"`)
              return renderedRoute
            },
            renderer: new Renderer({
              inject: {},
              headless: true,
              renderAfterElementExists: '#main',
          })
        }))
  }
}

module.exports = {
  baseUrl: '/',
  assetsDir: 'assets',
  filenameHashing: false,
  devServer: {
    disableHostCheck: true,
    proxyTable: {
      '/api': {
        logLevel: 'debug',
        target: 'http://0.0.0.0:8585',
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      },
      '/oauth': {
        logLevel: 'debug',
        target: 'http://0.0.0.0:8585',
        changeOrigin: true,
        pathRewrite: {
          '^/oauth': ''
        }
      }
    }
  }
};

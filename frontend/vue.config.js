module.exports = {
  assetsDir: 'assets',
  filenameHashing: false,
  devServer: {
    proxy: {
      '/api': {
        logLevel: 'debug',
        target: 'http://0.0.0.0:8585'
      },
      '/oauth': {
        logLevel: 'debug',
        target: 'http://0.0.0.0:8585/oauth/'
      }
    }
  }
};

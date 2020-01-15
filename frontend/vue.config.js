module.exports = {
  assetsDir: 'assets',
  devServer: {
    proxy: {
      '/api': {
        logLevel: 'debug',
        target: 'http://0.0.0.0:8585'
      }
    }
  }
};

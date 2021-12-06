'use strict';

const webpack              = require('webpack');
const merge                = require('webpack-merge');
const HtmlPlugin           = require('html-webpack-plugin');
const FriendlyErrorsPlugin = require('friendly-errors-webpack-plugin');
const helpers              = require('./helpers');
const commonConfig         = require('./webpack.config.common');
const environment          = require('./dev.env');
const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;


const webpackConfig = merge(commonConfig, {
  mode: 'development',
  devtool: 'inline-cheap-module-source-map',
  output: {
    path: helpers.root('dist'),
    publicPath: '/',
    filename: 'js/[name].bundle.js',
    chunkFilename: 'js/[name].chunk.js',
    devtoolModuleFilenameTemplate: '[absolute-resource-path]',
    devtoolFallbackModuleFilenameTemplate: '[absolute-resource-path]?[hash]'
  },
  optimization: {
    runtimeChunk: 'single',
    splitChunks: {
      chunks: 'all'
    }
  },
  plugins: [
    new webpack.EnvironmentPlugin(environment),
    new webpack.HotModuleReplacementPlugin(),
    new FriendlyErrorsPlugin(),
    new HtmlPlugin({
        template: 'public/index.html',
    }),
    new BundleAnalyzerPlugin({
      analyzerPort: 9090
    })
  ],
  devServer: {
    compress: true,
    historyApiFallback: true,
    hot: true,
    open: true,
    overlay: true,
    port: 8888,
    stats: {
      normal: true
    },
  proxy: {
      '/api': {
          logLevel: 'debug',
          target: 'http://0.0.0.0:8585'
      },
    '/scss': {
      logLevel: 'debug',
      target: 'http://0.0.0.0:8585'
    }
  }
  }
});

module.exports = webpackConfig;

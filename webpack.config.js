const path = require('path')
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const themePath = path.resolve(__dirname, 'themes/news.12bit.vn')

module.exports = {
  mode: process.env.NODE_ENV,
  entry: themePath + '/src/main.js',
  output: {
    filename: '[name].js',
    path: themePath + '/assets'
  },
  module: {
    rules: [
      {
        test: /\.scss$/,
        use: [
          process.env.NODE_ENV !== 'production' ? 'style-loader' : MiniCssExtractPlugin.loader,
          'css-loader',
          'sass-loader'
        ]
      }
    ]
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: 'theme.css',
      chunkFilename: '[id].css',
    })
  ]
}

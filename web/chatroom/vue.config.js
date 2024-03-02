const path = require('path');
const CompressionPlugin = require('compression-webpack-plugin')
const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  publicPath:'/static',
  transpileDependencies: true,
  devServer: {
    proxy: {
      //  /api是请求前缀
      // 代理服务器一
      '/v1': {//匹配所有以/阿皮开头的请求路径
        target: 'http://127.0.0.1:8082',//代理目标的基础路径
        pathRewrite:{

        },
        ws: true,//用于支持websocket 默认为真
        changeOrigin: true//用户控制请求头中的host值 默认为真
      },
    }
  },
  // configureWebpack: {
  //   plugins: [
  //     new CompressionPlugin({
  //       test: /\.(js|css)?$/i, // 哪些文件要压缩
  //       // filename: '[path].gz[query]',// 压缩后的文件名
  //       filename: '[path][base].gz',// 压缩后的文件名
  //       algorithm: 'gzip',// 使用gzip压缩
  //       threshold: 10240,
  //       minRatio: 0.8,// 压缩率小于0.8才会压缩
  //       deleteOriginalAssets: true // 删除未压缩的文件，谨慎设置，如果希望提供非gzip的资源，可不设置或者设置为false
  //     })
  //   ]
  // }
})

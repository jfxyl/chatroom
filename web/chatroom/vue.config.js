const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      //  /api是请求前缀
      // 代理服务器一
      '/v1': {//匹配所有以/阿皮开头的请求路径
        target: 'http://127.0.0.1:8080',//代理目标的基础路径
        pathRewrite:{

        },
        // ws: true,//用于支持websocket 默认为真
        // changeOrigin: true//用户控制请求头中的host值 默认为真
      },
    }

  },
})

const express=require('express');
const app=express();

// 配置解析表单数据的中间件
app.use(express.urlencoded({ extended: false }))

// 在配置cors中间件之前，配置JSONP的接口
app.get('/api/jsonp', (req, res) => {
    // 得到函数名
    const funcName = req.query.callback
    // 定义要发送到客户端的数据对象
    const data = {
        name: 'hang',
        age: 21
    }
    // 拼接出一个函数的调用
    const scriptStr = `${funcName}(${JSON.stringify(data)})`
    // 将拼接出的字符串响应给客户端
    res.send(scriptStr)
})

// 在路由之前，配置cors中间件，从而解决接口跨域问题
const cors = require('cors')
app.use(cors())

// 导入路由模块
const router = require('./router/router.js')
// 把路由模块注册到app上
app.use('/api', router)

// 启动服务器
app.listen(3000, () => {
    console.log('express server running at http://localhost:3000');
})
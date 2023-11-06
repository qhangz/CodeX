const express = require('express')
const router = express.Router()
const fs=require('fs')

// 用于post接收前端传过来的数据
// const bodyParser = require('body-parser')
// router.use(bodyParser.urlencoded({ extended: false }))
// router.use(bodyParser.json());

router.get('/test',(req,res)=>{
    const data=fs.readFileSync('./data/test.json')
    const testData = JSON.parse(data)
    res.send(testData)
})

// 挂载路由
router.get('/get', (req, res) => {
    console.log("/get");
    // 获取客户端发送到服务器的数据
    let query = req.query
    // 调用res.send()方法，将数据返回给客户端
    res.send({
        status: 200,
        msg: 'get请求成功',  // 状态描述
        data: query,  // 响应给客户端的数据
    })
})

// 定义post接口
router.post('/post', (req, res) => {
    // 获取客户端发送到服务器的数据
    let body = req.body
    // 调用res.send()方法，将数据返回给客户端
    res.send({
        status: 200,
        msg: 'post请求成功',  // 状态描述
        data: body,  // 响应给客户端的数据
    })
})

// 定义DELETE接口
router.delete('/delete', (req, res) => {
    // 获取客户端发送到服务器的数据
    let query = req.query
    // 调用res.send()方法，将数据返回给客户端
    res.send({
        status: 200,
        msg: 'delete请求成功',
        data: query
    })
})

// 定义PUT接口
router.put('/put', (req, res) => {
    // 获取客户端发送到服务器的数据
    let body = req.body
    // 调用res.send()方法，将数据返回给客户端
    res.send({
        status: 200,
        msg: 'put请求成功',
        data: body
    })
})


// 暴露路由
module.exports = router
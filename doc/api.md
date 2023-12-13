# api说明
```text
1. user:
login：http://localhost/api/user/login
register：/api/user/register
get user list: /api/user/list
user info: /api/user/info/:username

2. article：
hot article: /api/article/hot
follow artocle: /api/article/follow
frontend article: /api/article/frontend
backend article: /api/article/backend
get article by id: /api/article/info/:id

3. discuss
publish discuss: /api/discuss/publish
get discuss list: /api/discuss/list
get discuss by id: /api/discuss/info/:id
```
# 接口返回参数说明
1. user:
```json
// register
register sucess
{
    "code": "400",
    "msg": "register success"
}
register failed
{
    "code": "400",
    "error": "error message"
}

// login
login success
{
    "code": "200",
    "data": {
        "token": "1",
        "userInfo": {
            "age": 0,
            "avatar_image": "",
            "created_at": "2023-12-01T22:49:05+08:00",
            "summary": "",
            "username": "HANG"
        }
    },
    "msg": "登录成功"
}
login failed
{
    "code": "400",
    "error": "error message"
}

// user list
success
{
    "code": "200",
	"data": "userList",
}
failed
{
    "code": "400",
    "error": "error message"
}

// get user info by username
success
{
    "code": "200",
	"data": "userInfo",
}
failed
{
    "code": "400",
    "error": "error message"
}

// 返回的msg:
// register fail:
// register success ;StatusOK
// invalid registration information:数据验证失败
// user already exists
// email already exists
// email format is wrong 

// login fail:
// invalid login information
// user not exit
```

2. article:
```json
// article list
/api/article/hot
/api/article/follow
/api/article/frontend
/api/article/backend
{
    "code": "400",
    "msg": "success",
    "data": [
        {
            "id": 1,
            "author": "HANG",
            "title": "hello",
            "summary": "hello world",
            "createTime": "2019-01-01 12:00:00",
            "viewNum": 100,
            "likeNum": 100
        },
        {
            "id": 2,
            "author": "HANG",
            "title": "hello",
            "summary": "hello world",
            "createTime": "2019-01-01 12:00:00",
            "viewNum": 100,
            "likeNum": 100
        }
    ]
}

3. discuss:
```json
/api/discuss/publish (post)
{
    "code": "400",
    "msg": "success",
}
```
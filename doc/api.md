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
get discuss by id: /api/discuss/info
get the top title of discuss list(from pre to end): /api/discuss/toplist

4. comment
publish comment: /api/comment/publish
```
# 接口返回参数说明
```json
success
{
    "code": "200",
    "msg": "msg",
    "data": "data"
}
failed
{
    "coda": "400 or 500",
    "error": "error msg"
}
```
1. user:
```json
// register
register success
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

3. discuss
```json
// publish
success
{
    "coda": "200",
    "msg": "publish success"
}
// get discuss info
{
    "code": "200",
    "data": {
        "ID": 1,
        "CreatedAt": "2023-12-17T16:07:16.454+08:00",
        "UpdatedAt": "2023-12-17T16:07:16.454+08:00",
        "DeletedAt": null,
        "author": "HANG",
        "title": "frontend explore",
        "summaty": "codex explore",
        "content": "front end learning",
        "category": "frontend",
        "like_num": 17,
        "view_num": 77,
        "Comment": [
            {
                "ID": 1,
                "CreatedAt": "2023-12-20T19:04:31.296+08:00",
                "UpdatedAt": "2023-12-20T19:04:31.296+08:00",
                "DeletedAt": null,
                "DiscussID": 20,
                "author": "YOLO",
                "content": "快了",
                "like_num": 0,
                "view_num": 0
            },
        ]
    }
}
// get discuss list
{
    "code": "200",
    "data": [
        {
            "id": 8,
            "author": "HelloWorld",
            "title": "hello",
            "summary": "world",
            "like_num": 3,
            "view_num": 6,
            "created_at": "2023-12-17 16:12:12"
        },
    ]
}
// get top title of discuss
{
    "code": "200",
    "data": [
        {
            "id": 4,
            "title": "作业写完了吗"
        }
    ]
}
```

4. comment
```json
// publish
success
{
    "coda": "200",
    "msg": "publish success"
}
```
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
add discuss view number: /api/discuss/info/view
add discuss like number: /api/discuss/info/like
get discuss list by username: /api/discuss/minelist

4. comment
publish comment: /api/comment/publish
add comment view number: /api/comment/info/view
add comment like number: /api/comment/info/like
5. games
games list: /api/game/list

6. event
evnet list: /api/event/list
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
// get mine discuss
{
    "code": "200",
    "data": [
        {
            "id": 22,
            "author": "HANG",
            "title": "d3",
            "summary": "解决疑难杂症",
            "category": "frontend",
            "like_num": 0,
            "view_num": 9,
            "created_at": "2023-12-23 20:44:06"
        },
        {
            "id": 1,
            "author": "HANG",
            "title": "frontend explore",
            "summary": "codex explore",
            "category": "frontend",
            "like_num": 18,
            "view_num": 78,
            "created_at": "2023-12-17 16:07:16"
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
5. games
```json
// game list
success
{
    "coda": "200",
    "msg": "success",
    "data":{
        [
            "route": "/games/santaclaus",
            "banner_image": "",
            "game_name": ""
        ]
    }
}
```
6. event
```json
// game list
success
{
    "coda": "200",
    "msg": "success",
    "data":
    {
        "route": "",
        "url": "https://www.bagevent.com/event/8758229?bag_track=article1219",
        "name": "CodeX年度技术演讲",
        "time": "2024-01-17 周三",
        "place": "线上",
        "status": "未开始",
        "banner_image": "https://i.imgtg.com/2023/07/02/Okushr.jpg"
    },
}
```
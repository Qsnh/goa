# 用户接口

+ [用户分页列表](#用户分页列表)
+ [编辑用户](#编辑用户)
+ [用户详情](#用户详情)

### 用户分页列表

接口：`/backend/users`
方法：`get`

**请求：**

| 参数 | 说明 |
| --- | --- |
| `keywords` | 关键字，搜索昵称和邮箱 |

返回：

```
{
    "page": 1,
    "page_size": 10,
    "total": 4,
    "users": [
        {
            "id": 4,
            "nickname": "小滕哈哈哈",
            "avatar": "/static/uploads/avatar/28518dbd76acf56723b3900cefd29c31.jpg",
            "email": "13675626825@163.com",
            "is_lock": 1,
            "company": "",
            "age": 0,
            "profession": "",
            "website": "",
            "weibo": "",
            "wechat": "",
            "created_at": "2018-11-07T17:20:14+08:00",
            "updated_at": "2018-11-07T17:20:14+08:00",
            "questions": null,
            "answers": null
        },
        {
            "id": 3,
            "nickname": "小滕123",
            "avatar": "",
            "email": "12312312@qq.com",
            "is_lock": -1,
            "company": "",
            "age": 0,
            "profession": "",
            "website": "",
            "weibo": "",
            "wechat": "",
            "created_at": "2018-11-06T15:49:07+08:00",
            "updated_at": "2018-11-06T15:49:07+08:00",
            "questions": null,
            "answers": null
        }
    ]
}
```

### 编辑用户

接口：`/backend/user/:id`
方法：`get`

**请求：**

| 参数 | 说明 |
| --- | --- |
| `is_lock` | 是否锁定，1是，-1否 |
| `password` | 新密码，不修改请不要传 |

返回：

```
{
    "code": "0",
    "message": ""
}
```

### 用户详情

接口：`/backend/user/:id`
方法：`get`
返回：

```
{
    "user": {
        "id": 1,
        "nickname": "小滕",
        "avatar": "/static/uploads/avatar/e8e89f75ef964f6cb95f7c7edb88919e.jpg",
        "email": "12345@qq.com",
        "is_lock": 1,
        "company": "阿里云",
        "age": 23,
        "profession": "PHPER",
        "website": "https://v.58hualong.cn",
        "weibo": "xiaotenggogogo",
        "wechat": "wechat",
        "created_at": "0001-01-01T08:00:00+08:00",
        "updated_at": "0001-01-01T08:00:00+08:00",
        "questions": null,
        "answers": null
    }
}
```
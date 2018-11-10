# 回答接口

+ [回答列表](#回答列表)
+ [删除回答](#删除回答)

### 回答列表

接口：`/backend/answers`
方法：`get`

**请求：**

| 参数 | 说明 |
| --- | --- |
| `question_id` | 问题ID |

返回：

```
{
    "page": 1,
    "page_size": 10,
    "total": 22,
    "users": [
        {
            "User": {
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
            },
            "Question": {
                "User": {
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
                },
                "Category": {
                    "id": 2,
                    "name": "分类二",
                    "created_at": "0001-01-01T08:00:00+08:00",
                    "Questions": null
                },
                "id": 2,
                "title": "一步一步开发composer扩展包（三）：开始创建项目",
                "description": "## 二级标题",
                "view_num": 0,
                "is_ban": -1,
                "created_at": "2018-09-09T11:00:00+08:00",
                "updated_at": "2018-11-06T13:13:51+08:00",
                "answer_at": "2018-11-06T13:13:51+08:00",
                "answer_count": 2,
                "Answers": null,
                "AnswerUser": {
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
            },
            "id": 24,
            "content": "",
            "created_at": "2018-11-06T13:13:51+08:00",
            "updated_at": "2018-11-06T13:13:51+08:00"
        }
    ]
}
```

### 删除回答

接口：`/backend/answer/:id`
方法：`delete`
返回：

```
# 失败
{
    "code": "9",
    "message": "系统出错"
}

# 成功
{
    "code": "0",
    "message": ""
}
```
# 问题接口

+ [问题分页列表](#问题分页列表)
+ [问题删除](#问题删除)

### 问题分页列表

接口：`/backend/questions`
方法：`get`
返回：

```
{
    "page": 1,
    "page_size": 16,
    "questions": [
        {
            "User": {
                "id": 1,
                "nickname": "小滕",
                "avatar": "/static/uploads/avatar/e8e89f75ef964f6cb95f7c7edb88919e.jpg",
                "email": "12345@qq.com",
                "is_lock": -1,
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
                "is_lock": -1,
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
    ],
    "total": 4
}
```

### 问题删除

接口：`/backend/question/:id`
方法：`delete`
返回：

```
# 成功
{
    "code": "0",
    "message": ""
}

# 失败
{
    "code": "9",
    "message": "系统出错"
}
```
# 分类接口

+ [分类列表](#分类列表)
+ [分类信息](#分类信息)
+ [添加分类](#添加分类)
+ [编辑分类](#编辑分类)
+ [删除分类](#删除分类)

### 分类列表

接口：`/backend/categories`
方法：`get`
返回：

```
{
    "categories": [
        {
            "Id": 2,
            "Name": "分类二",
            "CreatedAt": "0001-01-01T08:00:00+08:00",
            "Questions": null
        },
        {
            "Id": 3,
            "Name": "分类三",
            "CreatedAt": "0001-01-01T08:00:00+08:00",
            "Questions": null
        }
    ]
}
```

### 分类信息

接口：`/backend/category/:id`
方法：`get`
返回：

```
{
    "category": {
        "Id": 2,
        "Name": "分类二",
        "CreatedAt": "0001-01-01T08:00:00+08:00",
        "Questions": null
    }
}
```

### 添加分类

接口：`/backend/category`
方法：`post`

**请求：**

| 参数 | 说明 |
| --- | --- |
| `name` | 分类名，不能为空，不能重复 |

返回：

```
# 失败
{
    "code": "5",
    "message": "不能为空"
}

# 成功
{
    "code": "0",
    "message": ""
}
```

### 编辑分类

接口：`/backend/category/:id`
方法：`put`

**请求：**

| 参数 | 说明 |
| --- | --- |
| `name` | 分类名，不能为空，不能重复 |

返回：

```
# 错误
{
    "code": "5",
    "message": "分类名已经存在"
}

# 成功
{
    "code": "0",
    "message": ""
}
```

### 删除分类

接口：`/backend/category/:id`
方法：`delete`

返回：

```
# 失败
{
    "code": "5",
    "message": "该分类下有问题，无法删除"
}

# 成功
{
    "code": "0",
    "message": ""
}
```

# 主面板

+ [简单的数据统计](#简单的数据统计)

### 简答的数据统计

接口：`/backend/dashboard`
方法：`get`

返回数据：

```
{
    "data": {
        "count_answer": 0,
        "count_question": 0,
        "count_register": 0
    }
}
```

解释：

| 字段 | 说明 |
| --- | --- |
| `count_answer` | 今日回答数量 |
| `count_question` | 今日提问数量 |
| `count_register` | 今日注册数量 |
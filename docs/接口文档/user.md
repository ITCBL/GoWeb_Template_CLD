# user 模块 

**接口文档示例**

## 1、登录

- url:  /api/v1/user/login
- method: POST
- headers:

| 字段 | 类型   | 说明                                   |
| ---- | ------ | -------------------------------------- |
| 无 |  |  |

- body-json

| 字段     | 类型   | 说明         |
| -------- | ------ | ------------ |
| username | String | 必填，用户名 |
| password | String | 必填，密码   |


- response示例:
```json
{
    "code": 0,
    "data": {
        "id": 1, // 用户ID
        "username": "root6",// 用户名
        "password": "",// 密码
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpX......KksTgR9c-7TTeGtc"// token
    },
    "error": null
}
```


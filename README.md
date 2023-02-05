# 上传

### Request

- Method:  **POST**

- URL: ```/upload```

- Body: 

  ```json
  {
    "token": "curtion",
    "files": [
      {
        "path": "1",
        "name": "1.txt",
        "content": "dGhpcyBpcyBhIGV4YW1wbGU="
      },
      {
        "path": "2",
        "name": "2.txt",
        "content": "MTIz"
      }
    ]
  }
  ```

### Response

```json
{
  "code": 200,
  "msg": "创建成功",
  "data": "success"
}
```

# 下载

### Request

- Method:  **GET**

- URL: ```/download?token=curtion```

| 参数  |        说明        | 是否必须 |
| :---: | :----------------: | :------: |
| token | Token |    是    |

### Response

```json
{
  "code": 200,
  "msg": "获取成功",
  "data": [
    {
      "path": "curtion\\1\\1.txt",
      "name": "1.txt",
      "content": "dGhpcyBpcyBhIGV4YW1wbGU="
    },
    {
      "path": "curtion\\2\\2.txt",
      "name": "2.txt",
      "content": "MTIz"
    }
  ]
}
```
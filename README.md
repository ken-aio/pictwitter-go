# Pictwitter

# API
## Fonts
フォントのリストを取得します  
```
$ curl -H "Content-Type: application/json" -X GET http://localhost:8888/api/v1/fonts | jq
```

```
[
  {
    "id": 2,
    "name": "test1",
    "download_url": "http://dummy",
    "created_at": "2016-09-09T00:00:00Z",
    "updated_at": "2016-09-09T10:00:00Z"
  },
  {
    "id": 3,
    "name": "test2",
    "download_url": "http://dummy2",
    "created_at": "2016-09-19T00:00:00Z",
    "updated_at": "2016-09-19T10:00:00Z"
  }
]
```

参考  
https://github.com/ken-aio/go-echo-base

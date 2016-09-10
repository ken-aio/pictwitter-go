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

## Images
言語を画像化したbase64を取得します  
```
$ curl -H "Content-Type: application/json" -X GET http://localhost:8888/api/v1/image?word=a
```

```
{"image":"data:image/jpeg;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAXCAAAAAAXEYEkAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfgCQoRNgQrqK8cAAAAe0lEQVQI12P8zwADTAxUZT5LUeZ1v8HAwMDw31161SFFtb///zP87zv8/383w+3//xn+////cLEbw67//5kY7rvLt39nYGFgYPiiqLz3/2yGw///sxy9P92J4SPDXwYGFgGGeWJXuxnWi2ky/KsVFgg+pCEw6z8jxU4HAADmLxYuWQRaAAAAJXRFWHRkYXRlOmNyZWF0ZQAyMDE2LTA5LTEwVDE3OjU0OjA0KzA5OjAwQOrTXgAAACV0RVh0ZGF0ZTptb2RpZnkAMjAxNi0wOS0xMFQxNzo1NDowNCswOTowMDG3a+IAAAAHdEVYdGxhYmVsAGGns5qFAAAAAElFTkSuQmCC"}
```

参考  
https://github.com/ken-aio/go-echo-base

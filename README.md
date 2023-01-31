# Logging and Echo server

This is a simple implementation of logging and echo server written in Go using [GinGonic](https://github.com/gin-gonic/gin).

I've copied `gin.logger` middleware to print out some information.

## Build and run
Easy run!
```shell
go build && ./logging-server -bind 0.0.0.0:8888
```

By default, if no `bind` paramenter is provided the server binds to `localhost:8080`. 

## Example
```shell
curl -vvv -d'{"key":"value","banana":"banana"}' -H 'Content-Type: application/json' localhost:8888/test
```
```
* Trying 127.0.0.1:8888...
* Connected to localhost (127.0.0.1) port 8888 (#0)
> POST /test HTTP/1.1
> Host: localhost:8888
> User-Agent: curl/7.87.0
> Accept: */*
> Content-Type: application/json
> Content-Length: 33
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Tue, 31 Jan 2023 10:42:55 GMT
< Content-Length: 43
< 
* Connection #0 to host localhost left intact
"{\"key\":\"value\",\"banana\":\"banana\"}"  
```

Default:
```
[GIN] 2023/01/31 - 11:42:55 | 200 |      31.851µs |       127.0.0.1 | POST     "/test"
```

Custom:
```
[GIN] 2023/01/31 - 11:42:55 | 200 |      31.851µs |       127.0.0.1 | POST     "/test"
Accept              : */*
Content-Length      : 33
Content-Type        : application/json
User-Agent          : curl/7.87.0

{"key":"value","banana":"banana"}

```

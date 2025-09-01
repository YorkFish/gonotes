## IM System Demo

- v1.0 基础 server 构建
- v2.0 用户上线功能
- v3.0 用户消息广播功能
- v4.0 用户业务封装
- v5.0 在线用户查询
- v6.0 修改用户名
- v7.0 超时强踢功能
- v8.0 私聊功能 `$ go build -o server main.go server.go user.go`
- v9 客户端
	- 9.1 建立连接 `$ go build -o client client.go`
	- 9.2 命令行解析 `$ ./client -h`
	- 9.3 菜单显示
	- 9.4 更新用户名
	- 9.5 公聊模式
	- 9.6 私聊模式

```
$ ./client -h
Usage of ./client:
  -ip string
        设置服务器IP地址（默认是 127.0.0.1） (default "127.0.0.1")
  -port int
        设置服务器端口地址（默认是 8888） (default 8888)
```

## Gin Demo

- v1.0

```
$ go mod init demo

$ go get github.com/gin-gonic/gin

GET  {host}/ping
POST {host}/ping/{id}
```

- v2.0

```
├── main.go
├── pojo
│   └── User.go
├── service
│   └── UserService.go
└── src
    └── UserRouter.go

add model User
add group router
```

- v3.0

```
DELETE {host}/v1/user/{id}
PUT    {host}/v1/user/{id}
```

- v4.0

```
create database `gintest`;

create table users (
    id int,
    name varchar(45),
    password varchar(45),
    email varchar(45),
    primary key(id)
);

insert into users (id, name, password, email) values
(1, 'york', '123456', 'example@gmail.com'),
(2, 'fish', '123456', 'example@gmail.com'),
(3, 'jessy', '123456', 'example@gmail.com');
```

```
GET
    get all
    get one
```

- v5.0

```
POST
PUT
DELETE
```

- v6.0

```
log

middleware

basic auth
router.Use(gin.BasicAuth(gin.Accounts{"Tom": "123456"}), middlewares.Logger())
```

- v7.0

```
Recovery
validator

go get github.com/go-playground/validator/v10
```

- v8.0

```
middleware
session

go get github.com/gin-contrib/sessions
go get github.com/gin-contrib/sessions/cookie
```


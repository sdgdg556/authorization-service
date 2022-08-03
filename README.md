# 身份认证鉴权服务demo

#### 项目简介

整个项目包括实现用户和角色创建以及删除、用户鉴权、用户角色校验和展示等功能。

允许不同用户绑定不同角色，通过获取auth_token来访问对应用户角色资源。

对外暴露一系列http接口来提供服务。

#### 运行环境

golang 版本>=1.13

#### 项目结构

#####     目录

- cmd: 项目入口，main.go内部会调用网关层的初始化

- configs：配置文件

- dao：数据存储层(本项目使用的非持久化内存进行存储)，承接服务层的请求，对数据进行处理

- helper: 一些公共方法，加密，随机字符串等

- model：对请求、响应、配置、存储对象封装成结构体的实例层

- server: 网关层，实现了自定义路由，请求解析等

- service: 服务层，承接网关层的请求，对请求以及响应进行处理

- test： 单元测试，主要包括所有涉及到的接口服务测试


#####     配置说明

configs下有两个配置文件：

- auth.json负责鉴权相关目前仅支持修改token过期时间，过期时间内部已经做了解析，支持"s"、"m"、"h"和数字的组合 例如默认配置"2h"代表token2小时过期。

- http.json负责httpclient配置，目前仅支持修改ip,port。


#### 接口列表

| 请求方式 | url | 接口描述 | 接口参数 | 参数示例 | 返回成功实例 |
| --- | --- | --- | --- | --- | --- |
| post | /user/create | 创建用户 | user_name (string, required) <br/>user_password(string, required) | {"user_name":"caohaoyu", "user_password": "abc"} | {"code":200,"message":"ok","data":null} |
| post | /user/delete | 删除用户 | user_name (string, required) <br/>user_password(string, required) | {"user_name":"caohaoyu", "user_password": "abc"} | {"code":200,"message":"ok","data":null} |
| post | /authorization | 用户授权 | user_name (string, required) <br/>user_password(string, required) | {"user_name":"caohaoyu", "user_password": "abc"} | {"code":200,"message":"ok","data":{"auth_token":"chqY7tCoZizKvPpMmzcG5fKlUcm6jw0QoiYg"}} |
| post | /invalidate | 删除令牌 | auth_token (string, required) | { "auth_token": "chqY7tCoZizKvPpMmzcG5fKlUcm6jw0QoiYg"} | {"code":200,"message":"ok","data":null} |
| post | /role/create | 创建角色 | role_name (string, required) | {"role_name": "super"} | {"code":200,"message":"ok","data":null} |
| post | /role/delete | 删除角色 | role_name (string, required) | {"role_name": "super"} | {"code":200,"message":"ok","data":null} |
| post | /user/add-role | 用户添加角色 | user_name (string, required) <br/>user_password(string, required)<br/>role_name (string, required) | {"role_name": "super","user_name": "caohaoyu","user_password": "abc"} | {"code":200,"message":"ok","data":null} |
| post | /user/check-role | 检查用户是否有某个角色 | auth_token (string, required)<br/>role_name (string, required) | {"auth_token": "4XsohHLtu95ZP92nbKtgQJsSLAJwyWcfP4z5","role_name": "hushi"} | {"code":200,"message":"ok","data":{"result":true}} |
| post | /user/roles | 用户下所有角色列表 | auth_token (string, required) | { "auth_token": "chqY7tCoZizKvPpMmzcG5fKlUcm6jw0QoiYg"} | {"code":200,"message":"ok","data":{"roles":["super"]}} |

#### 开始使用

#####     服务启动

拉取代码之后, 可以在goland里直接build main.go，如果是命令行里运行cd到项目根目录运行：

```
go run cmd/main.go       
```

看到运行日志输出start http server, listing...后即代表服务已经启动，可以使用postman模拟http请求访问api。

#####     单元测试

单元测试部分如果在命令行里运行可以cd到项目根目录运行：

```
go test -v test/service_test.go
```

在goland里可以直接到test目录下build service_test.go文件

#### 问题&&解决方案

- 如果出现bind: address already in use报错：
  解决：可以看看configs目录下的http.json配置，默认配置的8000端口可能由于本地8000端口已经被占用，可以修改port配置之后再重新启动服务。


#### 项目延伸

- 全局预留了context方便后续对请求做timeout处理。

- 目前基于内存存储且不涉及更新，存储实例都不带唯一id，如果涉及更新并且存储可持久化那么需要使用唯一id作为存储实例的主键。
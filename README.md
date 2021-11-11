# grpc_etcd_gin

这个demo对应文章 => [使用gin作网关, etcd作服务发现实现go-grpc微服务的保姆级教程](https://jayj1997.github.io/2021/11/09/gin_etcd_microservices/)

首先你需要开启etcd

然后修改源码中连接etcd的地址（除非你的也是:2379）

然后分别跑起ping和gateway

最后用postman或者bash工具

get请求一下<http://localhost:2379/v1/ping>通了没有即可

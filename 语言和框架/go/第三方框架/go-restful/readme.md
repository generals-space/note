参考文章

1. [Golang：go-restful库使用手册](https://blog.csdn.net/chszs/article/details/88974199)
2. [浅谈go-restful框架的使用和实现](https://www.jb51.net/article/137113.htm)

go-restful定义了Container WebService和Route三个重要数据结构。

Container 表示一个服务器，由多个WebService和一个 http.ServerMux 组成，使用RouteSelector进行分发
WebService 表示一个服务，由多个Route组成，他们共享同一个Root Path
Route 表示一条路由，包含 URL/HTTP method/输入输出类型/回调处理函数RouteFunction

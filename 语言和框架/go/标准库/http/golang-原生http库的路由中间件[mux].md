# golang-原生http库的路由中间件

1. [golang学习之negroni对于第三方中间件的使用分析](https://blog.csdn.net/kiloveyousmile/article/details/78740242)
2. [Golang构建HTTP服务（二）--- Handler，ServeMux与中间件](https://www.jianshu.com/p/16210100d43d)
3. [Go Web：自带的ServeMux multiplexer ](https://www.cnblogs.com/f-ck-need-u/p/10020942.html)
4. [Golang Web入门（2）：如何实现一个RESTful风格的路由](https://blog.csdn.net/inet_ygssoftware/article/details/117649919)
    - golang 内置的 mux 是无法实现 restful 接口的(因为 handler 没有对 GET/POST.. 做区分处理)

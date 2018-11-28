# golang-jwt token签发与认证

参考文章

1. [Go实战--golang中使用JWT(JSON Web Token)](https://blog.csdn.net/wangshubo1989/article/details/74529333)

2. [Golang构建HTTP服务（二）--- Handler，ServeMux与中间件](https://www.jianshu.com/p/16210100d43d)

3. [github - nan1888/beego_jwt/common/common.go](https://github.com/nan1888/beego_jwt/blob/e20949c6cb89e69310890a752d832929d1a1fc6b/common/common.go)

4. [基于golang从头开始构建基于docker的微服务实战笔记](https://blog.csdn.net/wdy_yx/article/details/79085588)

参考文章1中给出了完整的使用jwt对客户端登录请求签发token并返回, 之后的请求中使用中间件进行验证的示例. 不过由于ta引入了`negroni`路由中间件, 我认为没有必要, 然后我找到了参考文章2, 使用原生http的中方法实现中间件.

另外参考文章1中使用`jwt-go/request`直接对`http.Request`中携带的token进行验证, 考虑到以后可能使用beego, gin等各种第三方框架, 不可能直接得到`http.Request`对象, 按照参考文章3, 使用了`jwt.ParseWithClaims()`方法, 更为通用.
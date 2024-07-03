# 身份认证 session VS token

参考文章

1. [什么是 JWT -- JSON WEB TOKEN](https://www.jianshu.com/p/576dbf44b2ae)
2. [JSON Web Token 入门教程](http://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html)

我的认识: 

token相对于session, 是一种用CPU换内存的手段, 不需要在服务端保存session信息了.

关于权限, token中的确可以添加权限相关的字段, 访问接口时验证. 但必须要在数据库中保存用户权限, 才能在用户登录时, 签发token时创建其对应权限的字段.

至于每个接口如何验证用户请求是否具有访问权限, 还要依赖进一步权限归划. 这一点和使用session是相同的, 而和身份认证无关.

但是token虽然有过期时间, 但是服务端貌似没有类似能够强制用户下线的方法?

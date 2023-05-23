# SSO单点登陆及实现方式CAS OAuth

参考文章

1. [SSO单点登录原理详解（从入门到精通）](https://blog.csdn.net/qq1225308337/article/details/119040742)
    - 一文足够
    - SSO仅仅是一种设计架构, 而CAS和OAuth是SSO的一种实现方式, 他们之间是**抽象与具象的关系**. 
    - CAS流程
    - CAS实现的SSO, 不同域的服务, 它们之间的session不共享也是没问题的. 
2. [CAS、Oauth2还是SAML, 单点登录SSO方案该怎么选？](https://zhuanlan.zhihu.com/p/560667961)
    - OAuth2是一个授权协议, 主要用来作为API的保护, 我们称之为STS（安全令牌服务, Security Token Service）. 但是在某些情况下, 也可以被用来实现WEB SSO单点登录. 

参考文章1, 2对于 SSO 与 CAS 的解释都很清楚, 没什么有疑问的, 主要是 OAuth2.

参考文章2中说, **OAuth2是一个授权协议, 主要用来作为API的保护, 我们称之为STS（安全令牌服务, Security Token Service）. 但是在某些情况下, 也可以被用来实现WEB SSO单点登录**. 

就是说, ta本来并不是为 SSO 设计的, 不过可以用来实现.

其实典型的应用就是不同平台间的相互认证, 有很多技术类型网站(比如 csdn, gitee等), 都可以通过 github 账户登录, 或是媒体类型网站(如网易云音乐, 腾讯视频等), 可以使用QQ, 微信账户登录.

但实际上这些网站跟 github, QQ, 微信, 都不属于同一个平台, 甚至不是一家公司, 只不过ta们提供了账户服务, 因此可以变向实现 SSO 的能力.

![](https://gitee.com/generals-space/gitimg/raw/master/2023/dfda6a2886759ba49ddbf7c497f1a306.png)

将参考文章2中"基于OAuth 2.0单点登录"一节的配图, `SP`替换成"csdn", `IDP`替换成"github", 就不难理解了.

------

有一个问题, 就是人家 github 凭什么给你做账户管理, 凭什么给你授权?

好像是有一个授权申请的地方.

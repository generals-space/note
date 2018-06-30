# http接口测试

参考文章

1. [golang web开发 Handler测试利器httptest](https://www.jianshu.com/p/21571fe59ec4)

golang的`http/net/httptest`提供了http接口的测试方法, 可以在不启动http server的情况下直接测试handler函数. 让开发者可以直接编写request, 而不必再写`get`, `post`等客户端代码.

但是这个测试方法只适用于原生`net/http`库的代码, 对于目前各种web框架不能做到兼容, 所以目前不打算花时间学习这个.

...不过参考文章1讲解得挺不错的.
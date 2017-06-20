# Linux命令-curl

## 

参考文章

[在linux下使用curl访问 多参数url GET参数问题](http://blog.csdn.net/sunbiao0526/article/details/6831327)

假设url为`http://mywebsite.com/index.PHP?a=1&b=2&c=3`, web形式下访问url地址，使用`$_GET`是可以在后台获取到所有的参数

而在Linux下使用`curl http://mywebsite.com/index.php?a=1&b=2&c=3`, `$_GET`只能获取到参数`a`. 由于url中有`&`，其他参数获取不到，必须对&进行下转义才能`$_GET`获取到所有参数

`curl http://mywebsite.com/index.php?a=1\&b=2\&c=3`
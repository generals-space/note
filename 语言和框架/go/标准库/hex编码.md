参考文章

1. [golang标准库中的encoding/hex包](https://www.jianshu.com/p/859e04e2bdf2)

base64编码存在大小写和字符, 某些场景下可能只允许小写字母和数字, base64无法满足, base16可以, 但是golang没有加入到标准库, 可以使用hex代替.

# Python-读写文件

参考文章

1. [python读文件的三个方法read()、readline()、readlines()详解](http://blog.csdn.net/u010039733/article/details/47858189)
2. [【Python】python文件打开方式详解——a、a+、r+、w+、rb、rt区别](https://blog.csdn.net/ztf312/article/details/47259805/)
3. [Python-读写文件](https://www.cnblogs.com/jessicaxu/p/7679104.html)
4. [Python文件操作中的a，a+,w，w+，rb+,rw+,ra+几种方式的区别](https://blog.csdn.net/yang520java/article/details/82660786)

| 特性 | r    | w    | a    | r+   | w+   | a+   | rb   | wb   | rb+  | wb+  | ab+  |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 可读 | √    |      |      | √    | √    | √    |      |      |      |      |      |
| 可写 |      | √    |      | √    | √    | √    |      |      |      |      |      |
| 可读 |      |      |      |      |      |      |      |      |      |      |      |
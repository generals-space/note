# Python-读写文件

参考文章

1. [python读文件的三个方法read()、readline()、readlines()详解](http://blog.csdn.net/u010039733/article/details/47858189)
2. [【Python】python文件打开方式详解——a、a+、r+、w+、rb、rt区别](https://blog.csdn.net/ztf312/article/details/47259805/)
3. [Python-读写文件](https://www.cnblogs.com/jessicaxu/p/7679104.html)
4. [Python文件操作中的a，a+,w，w+，rb+,rw+,ra+几种方式的区别](https://blog.csdn.net/yang520java/article/details/82660786)
5. [ValueError: must have exactly one of create/read/write/append mode](https://www.cnblogs.com/wujily/p/12872926.html)
    - python 中文件打开操作的`mode`中没有`rw`
    - 合法的`mode`有: r、rb、r+、rb+、w、wb、w+、wb+、a、ab、a+、ab+

|                            | r    | w    | a    | r+      | w+   | a+   | rb   | wb   | rb+  | wb+  | ab+  |
| :------------------------- | :--- | :--- | :--- | :------ | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 可读                       | √    |      |      | √       | √    | √    |      |      |      |      |      |
| 可写                       |      | √    |      | √       | √    | √    |      |      |      |      |      |
| 无文件时创建               |      |      |      | ×(报错) | √    |      |      |      |      |      |      |
| 打开时清空                 |      | √    |      | √       | √    |      |      |      |      |      |      |
| 打开时不清空(写入时会追加) |      |      |      |         |      | √    |      |      |      |      |      |

# 笔记-Java

1.

关于类的构造方法与main方法，构造方法创建类对象，而main方法则只是充当一个入口程序。多文件工程中，主类中需要一个main方法，作为工程的入口方法，其他类中只需要有构造方法就行。

2.ArrayList与List

List是一个接口，ArrayList是一个类，ArrayList继承并实现了List，List并未有具体实现

ArrayList相当于动态数组(它与Vector并列)，C++里的vector，它的初始化形式：ArrayList<String> paths = new ArrayList<String>();

ArrayList<String> paths;只是声明, 并非初始化

3.Java中可以通过`System.go()`函数建议JVM回收对象, 但并没有办法强制回收.
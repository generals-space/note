# C++中Class与Struct的区别和联系

C中没有class的概念，struct只是一个包装数据的语法机制。

而在C++中，依然存在`struct`, 但它其实就是一种特殊的`class`。

从语法上，`class`和`struct`做类型定义时只有两点区别：

1. 默认继承权限。如果不明确指定，来自class的继承按照private继承处理，来自struct的继承按照public继承处理;

2. 成员的默认访问权限。class的成员默认是private权限，struct默认是public权限。

除了这两点，class和struct基本上就是一个东西。语法上也没有任何其他区别。

下面的说明可能有助于澄清一些直觉的关于struct和class的错误认知：

1. 都可以有成员函数，包括各类构造函数，析构函数，重载的运算符，友元类，友元结构，友元函数，虚函数，纯虚函数，静态函数等;

2. 都可以有一大`public`/`private`/`protected`修饰符在里边;

3. 都可以使用大括号的方式初始化(不过在C++中不提倡)，如A a = {1, 2, 3}。A可以是struct也可以是class，前提是这个类结构足够简单，比如所有成员都是public，所有成员都是简单类型，没有声明的构造函数，不是其他struct/class的子类等;

4. 都可以进行复杂的继承甚至多重继承，一个struct可以继承自一个class，反之亦可; b一个struct可以同时继承5个class和5个struct，虽然这样做不太好。

5. 如果class的设计需要注意OO的原则和风格，那struct的应用也需要注意这些问题。

最后，作为语言的两个关键字，除去定义类型时有上述区别之外，还有一点点：class还用于定义模板参数，就像typename，但struct不用于定义模板参数。
# C语言-const用法

## 1. 用const修饰函数参数

如果参数作输出用，不论它是什么数据类型，也不论它采用“指针传递”还是“引用传递”，都不能加const修饰，否则该参数将失去输出功能。const只能修饰输入参数.

### 1.1 如果输入参数采用“指针传递”，那么加const修饰可以防止意外地改动该指针，起到保护作用。

例如：

```
void StringCopy(char*strDestination, const char *strSource);
```

其中strSource是输入参数，strDestination是输出参数。给strSource加上const修饰后，如果函数体内的语句试图改动strSource的内容，编译器将指出错误。

### 1.2 如果输入参数采用“值传递”，由于函数将自动产生临时变量用于复制该参数，该输入参数本来就无需保护，所以不要加const修饰。

例如不要将函数void Func1(int x) 写成void Func1(const int x)。

同理不要将函数void Func2(A a) 写成void Func2(const A a)。其中A为用户自定义的数据类型。

对于非内部(类的实例对象的内部外部数据传递)数据类型的参数而言，象voidFunc(A a) 这样声明的函数注定效率比较低。因为函数体内将产生A类型的临时对象用于复制参数a，而临时对象的构造、复制、析构过程都将消耗时间。

为了提高效率，可以将函数声明改为voidFunc(A &a)，因为“引用传递”仅借用一下参数的别名而已，不需要产生临时对象。

但是函数voidFunc(A &a) 存在一个缺点：
"引用传递"有可能改变参数a，这是我们不期望的。解决这个问题很容易，加const修饰即可，因此函数最终成为void Func(const A &a)。
以此类推，是否应将void Func(int x) 改写为void Func(const int& x)，以便提高效率？完全没有必要，因为内部数据类型的参数(同在类实例对象中的数据，应该指的是成员函数的参数)不存在构造、析构的过程，而复制也非常快，"值传递"和"引用传递"的效率几乎相当。

### 1.3 "const&"修饰输入参数的用法总结
对于非内部数据类型的输入参数，应该将“值传递”的方式改为“const引用传递”，目的是提高效率。例如将voidFunc(A a) 改为voidFunc(const A &a)。
对于内部数据类型的输入参数，不要将"值传递"的方式改为"const引用传递"。否则既达不到提高效率的目的，又降低了函数的可理解性。例如voidFunc(int x) 不应该改为voidFunc(const int &x)。

## 2. 用const修饰函数返回值(形如const char *fn())

### 2.1 如果给使用"指针传递"方式的函数返回值加const修饰，那么函数返回值(即指针)的内容不能被修改，且该返回值只能被赋给加const修饰的同类型指针。

例如函数：

```
const char* GetString(void);
```

那么区分如下两种用法：

```
char* str = GetString();//编译错误
const char* str = GetString();//正确
```

### 2.2 如果函数返回值采用"值传递方式"，由于函数会把返回值复制到外部临时的存储单元中，加const修饰没有任何价值。

例如不要把函数int GetInt(void)写成const int GetInt(void);

同理不要把函数A GetA(void)写成const A GetA(void)，其中A为用户自定义的数据类型。

如果返回值不是内部数据类型，将函数A GetA(void)改写为connst A& GetA(void)的确能提高效率，但此时千万要小心，一定要弄清楚函数究竟是想返回一个对象的"拷贝"还是仅返回"别名"就可以了，否则程序会出错。

函数返回值采用"引用传递"的场合不多，一般只出现在类的赋值函数中，目的是为了实现链式表达。例如：

```
class A
{
   A& operate = (const A& other);//赋值函数
};
A a, b, c;//a, b, c为A的实例对象
a = b = c;//链式赋值
(a = b) = c;//不正常的链式赋值，但合法
```

如果将赋值函数的返回值加const修饰，那么该返回值的内容不允许被改动。上述代码中，语句a = b = c仍然正确，但语句(a = b) = c则是非法的。

## 3. const成员函数(形如void fnMember() const{...})

1.首先要明白

在普通的非const成员函数中，this的类型是一个指向类类型的const指针(相当于底层const)。即可以改变this所指向的值，但不能改变this所保存的地址。

在const成员函数中，this的类型是一个指向const类类型对象的const指针(顶层+底层const)。既不能改变this所指向的对象的值，也不能改变this所保存的地址。

任何不会修改数据成员的函数都应该声明为const类型。这样如果在编写const成员函数时，不慎修改了数据成员，或者调用了其它非const成员函数，编译器将指出错误，这无疑会提高程序的健壮性。以下程序中，类Stack的成员函数GetCount仅用于计数，从逻辑上讲GetCount应当为const函数。编译器将指出GetCount函数中的错误。

```
class Stack
{
    public:
        void Push(int elem);
        int Pop(void);
        int GetCount(void) const;//const成员函数
    private:
        int m_num;
        int m_data[100];
};
int Stack::GetCount(void) const
{
    ++ m_num;//编译错误，企图修改数据成员m_num
    Pop();//编译错误，企图调用非const函数
    return m_num;
}
```
2.关于const成员函数的几点规则：

a.const对象只能访问const成员函数，而非const对象可以访问任意的成员函数，包括const成员函数;
b.const对象的成员是不可修改的，然而const对象通过指针维护的对象却是可以修改的;
c.const成员函数不可以修改对象的数据，不管对象是否具有const性质。它在编译时，以是否修改成员数据为依据,进行检查;
d.然而加上mutable修饰符的数据成员，对于任何情况下通过任何手段都可修改，自然此时的const成员函数是可以修改它的。

对于上述规则中a的理解：可以想像，如果const成员函数fn1可以调用非const成员函数fn2，而fn2会修改类成员，就违背了将fn1声明为const成员函数的初衷。

3.const修饰成员函数的本质

```
Class A
{
    ...
    public:
        fn(int a);
    ...
}
```

上述代码中fn函数其实有两个参数，第一个是A* const this，另一个才是int类型的参数a。

如果不想fn函数改变参数a的值，可以把函数原型改为fn(const int a)，但如果我们不允许fn改变this指向的对象呢？因为this是隐含参数，const没有办法直接修饰它，就加在函数的后面了，fn(int a) const，表示this的类型是const A* const this。

const修饰`*this`是本质，至于说"表示该成员函数不会修改类对象的数据"之类的说法只是一个现象，根源在于*this是const类型的。


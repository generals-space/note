# C++拷贝构造函数理解

参考文章

[C++拷贝构造函数详解](http://blog.csdn.net/lwbeyond/article/details/6202256)

## 1. 拷贝构造函数概念

1 对于简单类型的对象来说，复制操作是很简单的，例如：

```
int a = 100;
int b = a;
```

而类对象与普通对象不同，类对象内部结构一般较为复杂，存在各种成员变量。

2 一个类对象拷贝的简单实例：

```c
#include <iostream>
using namespace std;
class Copy{
        private:
                int a;
        public:
                //construction
                Copy(int b); 
                //ordinary function
                void Show();
};
Copy::Copy(int b)
{
        this->a = b;
}
void Copy::Show()
{
        cout << this->a << endl;
}
int main()
{
        Copy copy1(100);
        //note:here is a operation of copy function,
        //not assignment in the initialization to the object copy2
        Copy copy2 = copy1;
        copy2.Show();

        return 0;
}
```

程序运行时输出100。

从上述代码中可以看出，系统为对象copy2分配了内存并完成了对对象copy2的复制过程。就类对象而言，相同类型的类对象是通过拷贝构造函数来完成整个复制过程的(虽然没有显示的定义)。

3 一个最简单的拷贝构造函数(对上述代码稍加改动)

```c
#include <iostream>
using namespace std;
class Copy{
        private:
                int a;
        public:
                //construction
                Copy(int b);
                //the copy construction
                Copy(const Copy& copyX);
                //ordinary function
                void Show();
};
Copy::Copy(int b)
{
        this->a = b;
}
Copy::Copy(const Copy& copyX)
{
        this->a = copyX.a;//you can make this other expression like "this->a = 2",have a try
}
void Copy::Show()
{
        cout << this->a << endl;
}
int main()
{
        Copy copy1(100);
        //note:here is a operation of copy function,
        //not assignment to initialize the object copy2
        Copy copy2 = copy1;//or "Copy copy2(copy1);" is also the same
        copy2.Show();

        return 0;
}
```

Copy(const Copy& copyX)是我们自行定义的拷贝构造函数。

可见，拷贝构造函数是一种**特殊的构造函数**，函数的名称必须和类名一致，它必须的一个参数是**本类型**的一个**引用变量**。

## 2. 拷贝构造函数的调用时机

### 2.1 对象以值传递方式传入函数参数

```c
#include <iostream>
using namespace std;
class Copy{
        private:
                int a;
        public:
                //construction
                Copy(int b);
                //the copy construction
                Copy(const Copy& copyX);
                //ordinary function
                ~Copy();
};
Copy::Copy(int b)
{
        this->a = b;
        cout << "create" << endl;
}
Copy::Copy(const Copy& copyX)
{
        this->a = copyX.a;
        cout << "copy" << endl;
}
Copy::~Copy()
{
        cout << "delete" << endl;
}
void getFunction(Copy copyX)
{
        cout << "test" << endl;
}
int main()
{
        Copy copy1(100);
        getFunction(copy1);
        return 0;
}
```

上述代码运行输出为：

```
create //the 34th line,call the construction automatically while creating the copy1 object
copy   //when the getFunction() get called,the copy1 was copy to copyX 
test   //the getFunction() get running
delete //the getFunction() completed,and the copyX was no longer living.It's time to run the deconstruction
delete //the main() was completed,and the copy1 was deconstructed
```

main中调用`getFunction()`时发生的事情步骤如下：

1. 当copy1对象传入形参时，会产生一个临时变量copyX;

2. 调用拷贝构造函数将copy1的值赋给`copyX`，`copyX`将存在于`getFunction()`的作用域中;

3. `getFunction()`执行完毕后，析构掉`copyX`对象。


### 2.2 对象以值传递的方式从函数返回

```c
#include <iostream>
using namespace std;
class Copy{
        private:
                int a;
        public:
                //construction
                Copy(int b); 
                //the copy construction
                Copy(const Copy& copyX);
                //ordinary function
                ~Copy();
};
Copy::Copy(int b)
{
        this->a = b;
        cout << "create" << endl;
}
Copy::Copy(const Copy& copyX)
{
        this->a = copyX.a;
        cout << "copy" << endl;
}
Copy::~Copy()
{
        cout << "delete" << endl;
}
Copy getFunction()
{
        Copy copy1(2);
        return copy1;
}
int main()
{
        getFunction();
        return 0;
}
```

运行结果为：

```
create
delete
```

解释：上述结果是在linux下用g++进行编译，看起来似乎并没有调用拷贝构造函数，不过在VC下是正常调用的，应该是和编译器有关。

当`getFunction()`执行到`return`时，经历的如下步骤：

1. 产生一个临时变量tmp

2. 调用拷贝构造函数将copy1的值赋给tmp

3. 在函数执行到最后先析构copy1局部变量

4. 等`getFunction()`执行完毕再析构tmp对象

### 2.3 对象需要通过另外一个对象进行初始化：

```c
Copy copy1(100);
Copy copy2 = copy1;
//Copy Copy2(copy1);
```

后面两句都会调用拷贝构造函数。

## 3. 深浅拷贝

### 3.1 默认构造拷贝函数

很多时候在我们都不知道拷贝构造函数的情况下，传递对象给函数参数或者函数返回对象都能很好的进行，这是因为编译器会给我们自动产生一个"默认拷贝构造函数"，这个构造函数很简单，仅仅使用"老对象"的数据成员的值对"新对象"的数据成员一一赋值，它一般具有以下形式：

```c
Rect::Rect(const Rect& C)
{
    width = C.width;
    height = C.height;
}
```

默认构造函数不需要自行编写，编译器会自动生成。但是有些情况是默认构造函数无法处理的。例如：

```c
#include <iostream>
using namespace std;
class Rect
{
        private:
                int width;
                int height;
                static int count;
        public:
                Rect();
                ~Rect();
                static int getCount();
};
Rect::Rect()
{
        count ++; 
}
Rect::~Rect()
{
        count --; 
        //output the count value while deconstructing
        cout << "When deconstructed, the count is : " << count << endl;
}
int Rect::getCount()
{
        return count;
}
int Rect::count = 0;
int main()
{
        Rect rect1;
        cout << "The count of Rect is : " << Rect::getCount() << endl;

        Rect rect2 = rect1;
        cout << "The count of Rect is : " << rect2.getCount() << endl;
        return 0;
}
```

上述代码运行结果为：

```
The count of Rect is : 1
The count of Rect is : 1
When deconstructed, the count is : 0
When deconstructed, the count is : -1 //It means that there is only one object and the deconstruction was called twice
```

`class Rect`中有一个静态成员count，目的是为了计数(**对象的个数**)。main函数中创建了rect1对象，并由此复制出对象rect2，再输出count，按照理解，此时应该有两个对象存在，但实际程序运行时，输出的都是1，得到的只有一个对象，**这是不合理的**。因为，在销毁对象时，类的析构函数执行了两次(所以count才会为-1)，所以是销毁了两个对象。这样的话，唯一的解释就是，**拷贝构造函数没有正确(按照我们的想法)的处理静态成员**。

我们想要的是，在复制对象时，静态成员变量也需要改变。至于如何变，就需要自己按照需要来设计代码了。依照上面的要求，我们需要自行编写一个拷贝构造函数，示例如下：

```c
Rect::Rect(const Rect& r)
{
    width = r.width;
    height = r.height;
    count ++;//we need to add the count by ourselves
}
```

### 3.2 浅拷贝

所谓浅拷贝，指的是在对象复制时，只对对象中的数据成员进行简单的赋值，默认拷贝构造函数执行的也是浅拷贝。大多情况下"浅拷贝"已经能很好的工作了，但一旦对象存在了动态成员(例如new创建的对象)，那么浅拷贝就会出问题的。示例如下：

```c
#include <iostream>
using namespace std;
class Rect
{
        private:
                int width;
                int height;
                int *p; 
        public:
                Rect();
                ~Rect();
};
Rect::Rect()
{
        this->p = new int[100]; //this will make p point to an address in heap
}
Rect::~Rect()
{
        if(p != NULL)
                delete p;
}
int main()
{
        Rect rect1;
        Rect rect2 = rect1;
        return 0;
}
```

在上述代码运行结束之前，会出现一个运行错误。原因在于进行对象复制时，对于动态分配的内容没有进行正确的操作。分析如下：

在声明`rect1`对象后，由于在构造函数中有一个动态分配的语句，因此执行捕捞内存情况大致如下：

![](https://gitimg.generals.space/b8256f919f08dc2fe204972c9b666e88.jpg)

在使用`rect1`复制`rect2`时，由于执行的是浅拷贝，只是将成员的值进行赋值，这时`rect1.p = rect2.p`，即这个指针指向了堆里的同一个空间，如图：

![](https://gitimg.generals.space/877776a98ac0a9228727188c0dacc303.jpg)

这不是我们期望的效果，而且在销毁对象时，两个对象的析构函数将对同一个内存空间释放两次，这就是出错的原因。我们需要的不是两个p有相同的值，而是两个p指向的空间有相同的值，解决办法就是使用"深拷贝"。

### 3.3 深拷贝

在"深拷贝"情况下，对于对象中动态成员，就不能仅仅简单的赋值了，而应该重新动态分配空间。方法如下：

```
class Rect
{
        private:
                int width;
                int height;
                int *p;
        public:
                Rect();
                Rect(const Rect& r);
                ~Rect();
};
Rect::Rect()
{
        this->p = new int[100];
}
Rect::Rect(const Rect& r)
{
        this->width = r.width;
        this->height = r.height;
        this->p = new int;
        *(this->p) = *(r.p);
}
Rect::~Rect()
{
        if(p != NULL)
                delete p;
}
```

这样，在完成对象的复制后，内存的大致情况如下：

![](https://gitimg.generals.space/c5bab434b161318ef09f25d89f7bca20.jpg)

## 4. 阻止默认拷贝发生(不知道有啥用)

通过对对象复制的分析，我们发现对象的复制大多在进行"值传递"时发生，这里有一个小技艺可以防止按值传递--声明一个私有的拷贝构造函数。甚至不必去定义这个拷贝构造函数，这样因为拷贝构造函数是私有的，如果用户试图按值传递或函数返回该类对象，将得到一个编译错误，从而可以避免按值传递或返回对象类型。

## 5. 拓展

### 5.1. 关于上面那段浅拷贝代码的错误

浅拷贝"this->p = r.p"会导致三个错误：

1. this->p原来指向的内存空间未被释放，造成内存泄露;

2. 两个p指针指向同一块内存，任何一方变动都会影响另一方;

3. 在对象被析构时，p所指向的内存将会被释放两次，会产生严重错误。

### 5.2 为什么拷贝构造函数的参数只能使用引用类型

普通拷贝构造函数的应用：

```c
Rect rect1;
Rect rect2(rect1);
```

拷贝过程实际上相当于`rect2.Rect(rect1)` (不管rect1是否为引用类型)**不过这是无法写出来的**。

假设拷贝构造函数的参数不是引用类型而是类类型的话，就是按照"值传递"的方式传参，就会变成`rect2.Rect(Rect rectA)`，其中rectA是形参。但首先就应想到，要将rect1的值赋值给rectA，这样又要用到拷贝构造函数，`rectA.Rect(Rect rectB)`，还会有`rectC`，`rectD`...这样永远递归下去。


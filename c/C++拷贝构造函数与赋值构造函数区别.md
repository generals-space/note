# C++拷贝构造函数与赋值构造函数区别

拷贝构造函数与赋值构造函数相似，需要记住几点他们的区别

拷贝构造函数是在对象创建时调用来进行初始化的;赋值构造函数是只能被已经存在的对象调用。

赋值构造函数也有深拷贝和浅拷贝的问题。

赋值构造函数参数和返回值类型必须是引用类型。

重载运算符本质上是函数，其函数名为"operator="。

赋值构造函数实例：

```c
#include <iostream>
using namespace std;
class Assign
{
        private:
                int test;
        public:
                Assign(int test);
                Assign(const Assign& A); 
                ~Assign(){};
                Assign& operator=(const Assign& A); 
};
Assign::Assign(int test):test(test)
{
        cout << "create" << endl;
}
Assign::Assign(const Assign& A)
{
        this->test = A.test;
        cout << "copy" << endl;
}
Assign& Assign::operator=(const Assign& A)
{
        this->test = A.test;
        cout << "assignment" << endl;
        return *this;
}
int main()
{
        Assign assign1(1);
        Assign assign2(assign1);
        Assign assign3(3);
        assign3 = assign2;
        return 0;
}
```

上述代码运行结果为：

```
create //the 30th line calls the construction for assign1
copy   //the 31th line calls the copy construction for assign2
create //the 32th line calls the construction for assign3 
assignment //the 33th line calls the assignment construction for assign3
```
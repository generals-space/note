# 静态成员函数中使用非静态成员变量

静态成员函数不能访问非静态成员变量/函数，因为静态函数属于整个类而不是某个类对象，它不包含this指针，没有办法访问类对象的成员。

除了将变成员量改为静态成员变量，还有一种方法可以让静态成员函数访问本来是非静态的成员变量。

```c
#include <iostream>
using namespace std;
class Base
{
        private:
                int age;
        public:
                Base(int _age);
                void  showAge();
                static void changeAge(Base* _base, int _newAge);
};
Base::Base(int _age)
{
        this->age = _age;
}
void Base::showAge()
{
        cout << this->age << endl;
}
void Base::changeAge(Base* _base, int _newAge)
{
        _base->age = _newAge;
}
int main()
{
        Base base(12);
        base.showAge();
        Base::changeAge(&base, 14);
        base.showAge();
        return 0;
}
```

这样调用类静态成员函数的前提是必须实例化出实际对象，因为Base::changeAge()需要一个对象指针来充当this指针的角色。

当然，changeAge()的参数也可以是对象的引用。

注意，如果直接使用传值调用，编译可以成功，但结果不会正确。因为如果是传值调用，Base::changeAge()会拷贝一个新的newBase对象，并将其age成员变量改为14，然后函数执行完毕，这个newBase对象生命周期结束被销毁，而原来的base对象base成员的值并未发生改变。
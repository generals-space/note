普通成员方法与私有方法不能被类方法调用, 起码通过cls.__私有方法不行, 只能使用类的实例对象调用.

```
TypeError: unbound method validate_host_ip() must be called with Validator instance as first argument (got list instance instead)
```
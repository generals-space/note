# Python-local nonlocal global局部变量与全局变量

参考文章

1. [nonlocal和global的区别](https://blog.csdn.net/lyon____/article/details/118387002)

```py
def isMatch():
    now_state_set = {0}

    def update_now_state_set():
        nonlocal now_state_set
        now_state_set = {1}
    update_now_state_set()
```

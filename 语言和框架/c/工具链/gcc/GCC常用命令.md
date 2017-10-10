# GCC常用命令

生成a.out文件, 可执行.

```
$ gcc [-g] test.c [-o 可执行文件名]
```

生成一个.o中间文件

```
$ gcc -c test.c [-o 中间文件名]
```

生成.so共享库文件.

```
gcc -fPIC -c test.c
gcc -shared test.c -o test.so
``` 

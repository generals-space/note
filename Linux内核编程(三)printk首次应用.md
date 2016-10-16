# Linux内核编程(三)printk首次应用

`printk`是内核打印函数，需要在Linux内核源代码中调用printk函数需要包含`kernel.h`(`/include/linux/kernel.h`)

内核入口函数`start_kernel`(位于`/init/main.c`)首行添加代码：


```
printk(KERN_NOTICE "I'm the printk in the kerneln");
```

重新编译内核后，察看dmesg的信息，在开头处可以看到

```
[ 0.000000] I'm the printk in the kernel

[ 0.000000] Linux version 3.2.28 (root@leomass-virtual-machine) (gcc version 4.6.3 (Ubuntu/Linaro 4.6.3-1ubuntu5) ) #1 SMP Tue Nov 6 12:47:01 CST 2012
```

可以看到我们自己插入的printk执行成功了，在源代码的其他地方添加printk就可以打印想要跟踪的数据了。

注：`dmesg`输出的信息大小有限制，应该不能超过`ring buffer`的值，不然开始时输出的信息会被覆盖
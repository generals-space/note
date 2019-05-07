# golang环境变量-GOGC垃圾回收控制

参考文章

1. [Go 语言运行时环境变量快速导览](https://blog.csdn.net/htyu_0203_39/article/details/50852856)
 
`GOGC` 是Go Runtime最早支持的环境变量, 甚至比`GOROOT`还早, 几乎无人不知. 

`GOGC` 用于控制GC的处发频率, 其值默认为`100`, 意为直到自上次垃圾回收后`heap size`已经增长了100%时GC才触发运行. 即是`GOGC=100`意味着live heap size 每增长一倍, GC触发运行一次. 

如设定`GOGC=200`, 则live heap size 自上次垃圾回收后, 增长2倍时, GC触发运行. 总之, 其值越大则GC触发运行频率越低,  反之则越高. 

如果`GOGC=off` 则关闭GC.

虽然go 1.5引入了低延迟的GC, 但是`GOGC`对GC运行频率的影响不变. 

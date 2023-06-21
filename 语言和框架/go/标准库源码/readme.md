参考文章

1. [6. 开篇《 刻意学习 Golang - 标准库源码分析 》](https://learnku.com/articles/25470)
    - 所有Package集合简介

2. [Go 语言编译器的 "//go:" 详解](https://www.jianshu.com/p/afd6dd988c20)
    - `//go:`等价于C语言中的`#include`, 是给编译器看的标记
    - 常用标记: `//go:noinline`: 不要内联; `//go:nosplit`: 跳过栈溢出检测; `//go:noescape`: 禁止逃逸; `//go:norace`: 跳过竞态检测;

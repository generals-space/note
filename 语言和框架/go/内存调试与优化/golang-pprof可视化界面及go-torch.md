# golang-pprof可视化界面及go-torch

参考文章

1. [Golang性能调优(go-torch, go tool pprof)](https://blog.csdn.net/WaltonWang/article/details/54019891)

有关pprof的可视化展现方式有两种:

1. `go tool pprof`命令行中的web/svg子命令, 依赖[Graphviz](http://www.graphviz.org/)

2. [Uber go-torch](https://github.com/uber/go-torch), 依赖[FlameGraph](https://github.com/brendangregg/FlameGraph.git)

参考文章1种介绍了这两种方式的初级使用方法.

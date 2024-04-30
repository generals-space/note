# go mod vendor - modules.txt变更本地依赖[tidy]

参考文章

1. [How to remove an installed package using go modules](https://stackoverflow.com/questions/57186705/how-to-remove-an-installed-package-using-go-modules)

## 问题描述

一个 go mod 项目, 使用`go mod vendor`会把所依赖的包移到工程目录的 vendor 子目录中.

但是需要删除某个依赖包时, 从 go.mod 文件中删除, 但是再执行`go mod vendor`是不会把对应的依赖从 vendor 子目录中删除的.

## 解决方案

先执行`go mod tidy`, 再执行`go mod vendor`就可以了.

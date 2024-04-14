# git update-index --assume-unchanged忽略指定文件的变更[gitignore]

参考文章

1. [Git忽略部分修改的方法（.gitignore添加忽略文件不起作用的解决办法）](https://www.cnblogs.com/gongxianjin/p/17510858.html)

我们知道`.gitignore`是针对未加入git仓库的文件进行忽略的机制, 就是说在一个/一类文件还未加入仓库前, 在`.gitignore`文件中编写好相应的规则, 就可以忽略这类文件的"新增"行为, 避免误将其加入仓库.

但是对于已经存在于仓库中的文件, 再在`.gitignore`文件上编写规则, 是无法忽略开发者对其的"更新"行为的.

最近遇到了一个场景, kubebuilder 工程中从另一个外部项目拷贝了几个 crd 的`_type.go`文件, 对应的 yaml 也放到了`config/crd/bases`目录. 

但是在工程根目录下执行`make`重新生成 yaml 时, 总是会同时更新这几个外部 crd 的 yaml 文件, 这是我们不希望看到的, 我们希望外部的 crd yaml 老老实实待着, 如果外部项目发生变化(频率不高), 由我们手动更新.

参考文章1中介绍了一种方案, 如下:

添加忽略

```
git update-index --assume-unchanged ./config/config.yaml
```

取消忽略

```
git update-index --no-assume-unchanged ./config/config.yaml
```

有效.

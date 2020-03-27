# non-name xyz on left side of

参考文章

1. [Golang常见误区(一)](https://www.cnblogs.com/zivli/p/10214729.html)
2. [Struct field, error and non-name xyz on left side of :=](https://golang-nuts.narkive.com/OKOKbFOy/struct-field-error-and-non-name-xyz-on-left-side-of)
3. [Short variable declarations](https://golang.org/ref/spec#Short_variable_declarations)

主调函数

```go
type Configuration struct {
	KubeConfigFile string
	KubeClient     kubernetes.Interface
}

func caller(){
    config := &Configuration{
		KubeConfigFile: *argKubeConfigFile,
	}
	config.KubeClient, err := kubeclient.NewKubeClient(config.KubeConfigFile)
	if err != nil {
		return nil, err
	}
}
```

被调函数

```go
func NewKubeClient(kubeConfigFilePath string) (client *kubernetes.Clientset, err error) {
    // ...省略
	client, err = kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Errorf("init kubernetes client failed %v", err)
		return nil, err
	}
	return
}

```

在编译时, 主调函数在调用行出现了如下报错.

```
non-name config.KubeClient on left side of :=
```

理论上, 由于`config.KubeClient`为`interface`类型, 应该可以接收`NewKubeClient()`返回的`*kubernetes.Clientset`的结果.

网上大部分文章都只说了解决方法, 没说理论依据. 

由于在主调函数中`config.KubeClient`已经是一个已知变量, 可以同样把`err`也预先声明成已知变量.

```go
    var err error
	config.KubeClient, err = kubeclient.NewKubeClient(config.KubeConfigFile)
```

我本来以为修改`NewKubeClient()`的函数原型, 把返回值client类型也变成`interface`类型, 主调函数不变, 也可以成功(认为`:=`无法在赋值的过程中推导左侧变量类型, 不如直接把右侧返回值改成和左侧的一致.).

```go
func NewKubeClient(kubeConfigFilePath string) (client kubernetes.Interface, err error) {
}
```

但是失败了.

关于结构体成员不能使用`:=`进程赋值的原因, 可以见参考文章3, 官方结出的文档.

短型变量声明`:=`本质上是常规变量声明`var str string`的语法糖, golang在编译`:=`的时候会把左侧的变量按照右侧函数的返回值类型, 进行隐式的预声明(为了不与可能的同名变量冲突, 隐式声明只发生在`:=`当前所在的`{}`块中).

但由于已知(这里的已知表示是已经声明过的)的结构体成员也算是已知的, 所以不能把结构体成员和`:=`一起使用.

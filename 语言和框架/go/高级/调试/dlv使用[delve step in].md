# dlv使用[delve step in]

- dlv: 1.7.1
- golang: 1.15.15

试用了下dlv, 感觉和gdb还是有些区别.

dlv不必要求工程在GOPATH目录下, 就按平常的开发路径就可以.

`dlv debug ./main.go`可以直接从源码启动类似gdb的界面, `dlv exec main`可以调试二进制文件.

进入交互式界面后, 直接执行`l(ist)`指令, 会得到如下错误

```
(dlv) ls
Stopped at: 0x7f43ee6a4140
=>   1:	no source available
```

这一点和gdb不同, 在dlv中, 需要先执行`b main.main`打个断点, 然后执行`c(ontinue)`开始运行程序, 然后才能看到文件源码...

```console
$ dlv exec ./main
Type 'help' for list of commands.
(dlv) ls
Stopped at: 0x7fd2ba301140
=>   1:	no source available
(dlv) break main.main
Breakpoint 1 set at 0x1251e1b for main.main() ./main.go:43
(dlv) c
> main.main() ./main.go:43 (hits goroutine(1):1 total:1) (PC: 0x1251e1b)
Warning: debugging optimized function
    38:	func onDelete(obj interface{}) {
    39:		pod := obj.(*corev1.Pod)
    40:		klog.Infof("delete a pod: +v", pod.Name)
    41:	}
    42:
=>  43:	func main() {
    44:		home := homedir.HomeDir()
    45:		kubeconfig := filepath.Join(home, ".kube", "config")
    46:		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    47:		if err != nil {
    48:			panic(err)
(dlv) ls
```

且`l(ist)`命令不再能上下移动了, 只能显示当前执行到的部分.

## 不能对`go func()`执行`step in`

idea 与 vscode 的调试都是在底层调用dlv, 但是ta们两个都没有办法对`go func()`语句执行`step in`, 单步步进会直接跳过.

但是在直接使用dlv进入交互式命令行, 发现还是这样...

```go
(dlv) n
> k8s.io/client-go/informers.(*sharedInformerFactory).Start() /usr/local/gopath/pkg/mod/k8s.io/client-go@v0.17.2/informers/factory.go:137 (PC: 0x11f975e)
Warning: debugging optimized function
   132:		f.lock.Lock()
   133:		defer f.lock.Unlock()
   134:
   135:		for informerType, informer := range f.informers {
   136:			if !f.startedInformers[informerType] {
=> 137:				go informer.Run(stopCh)
   138:				f.startedInformers[informerType] = true
   139:			}
   140:		}
   141:	}
   142:
(dlv) s
> k8s.io/client-go/informers.(*sharedInformerFactory).Start() /usr/local/gopath/pkg/mod/k8s.io/client-go@v0.17.2/informers/factory.go:138 (PC: 0x11f9791)
Warning: debugging optimized function
   133:		defer f.lock.Unlock()
   134:
   135:		for informerType, informer := range f.informers {
   136:			if !f.startedInformers[informerType] {
   137:				go informer.Run(stopCh)
=> 138:				f.startedInformers[informerType] = true
   139:			}
   140:		}
   141:	}
   142:
```

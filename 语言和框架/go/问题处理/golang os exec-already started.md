# golang os exec-already started

参考文章

1. [Go os/exec Command.Start() twice in a row](https://stackoverflow.com/questions/39239499/go-os-exec-command-start-twice-in-a-row)

golang: 1.12

```go
	fmt.Println("running")
	cmd := exec.Command("/bin/bash", "./bash.sh")
	cmd.Run()
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("command error: %s\n", err) // command error: exec: already started
		return
	}
	fmt.Println(output)
```

执行上面代码时出现`exec: already started`问题, 原来`Output()`方法不只是获取命令的执行结果, 而是先调用该命令再获取其结果. 而 `exec.Command()` 得到的 cmd 对象, `Start()`, `Run()`, 及 `Output()`方法不可同时使用. 如果使用了 Run() 方法, 又想获取 ta 的打印结果, 可以将 cmd.Stdout 定位到某一个文件.


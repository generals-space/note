# golang-获取主调函数名

参考文章

1. [如何在Go的函数中得到调用者函数名?](https://colobu.com/2018/11/03/get-function-name-in-go/)
    - runtime.Caller

Caller打印当前协程调用栈中相关函数所在的文件名, 行号等信息.

参数 skip 表示调用过程中的函数栈帧数(从下往上计数), 0表示当前函数(调用`Caller(0)`的函数).

(由于历史原因, `Caller`与`Callers`的skip参数含义并不相同).

Caller的返回值包括:

1. pc程序计数器
2. file函数所在文件名(绝对路径))
3. line行号(函数定义处的行号).

返回值中的ok与 `if _, ok := map["string"]; ok {}`中的ok含义相同, 如果为false则无其他值.

```go
func main() {
	foo()
}
func foo() {
	// This is: main.foo, pc: 17385455, file: /Users/general/Code/playground/go-caller/main.go, line: 13, ok: bool
	printFuncInfo()
	bar()
}
func bar() {
	// This is: main.bar, pc: 17385461, file: /Users/general/Code/playground/go-caller/main.go, line: 18, ok: bool
	printFuncInfo()
}

func printFuncInfo() {
	// Caller() 的参数, 0表示当前函数, 1表示第1层主调函数, 第n层就是第n层主调函数
	pc, file, line, ok := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Printf(
		"This is: %s, pc: %d, file: %s, line: %d, ok: %T\n",
		funcName, pc, file, line, ok,
	)
}
```

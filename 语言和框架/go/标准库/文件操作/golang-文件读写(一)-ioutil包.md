# golang-文件读写(一)-ioutil包

参考文章

1. [Go语言学习之ioutil包(The way to go)](https://blog.csdn.net/wangshubo1989/article/details/69395568)

文件读写对我来说一直是一个很复杂的事情, 路径获取(当前路径, 相对路径与绝对路径), 判断文件存在及类型, 读取文件内容, 读取方式等...

golang里的文件读写操作是很多的, 入手的话通过ioutil包比较方便, 因为太简单了.

go v1.10.3中, ioutil只提供了如下几个方法

1. `ReadAll func(r io.Reader) ([]byte, error)`

2. `ReadDir func(dirname string) ([]os.FileInfo, error)`

3. `ReadFile func(filename string) ([]byte, error)`

4. `TempDir func(dir, prefix string) (name string, err error)`

5. `TempFile func(dir, prefix string) (f *os.File, err error)`

6. `WriteFile func(filename string, data []byte, perm os.FileMode) error`

7. `NopCloser func(r io.Reader) io.ReadCloser`

## 1. ReadAll

```go
ReadAll func(r io.Reader) ([]byte, error)
```

读取 r 中的**所有数据**, 返回读取的数据和遇到的错误.  

如果读取成功, 则 err 返回 nil, 而不是 EOF, 因为 ReadAll 定义为读取所有数据, 所以不会把 EOF 当做错误处理. 

`r`是一个`io.Reader`, 这是一个接口类型. 比如`net/http`包中http响应的`Body`对象.

## 2. ReadDir

```go
ReadDir func(dirname string) ([]os.FileInfo, error)
```

读取指定目录中的所有目录和文件（不包括子目录）.  

返回读取到的文件信息列表和遇到的错误, 列表是经过排序的. 

`dirname`必须是一个目录路径, 否则会出错.

## 3. ReadFile

```go
ReadFile func(filename string) ([]byte, error)
```

读取文件中的所有数据, 返回读取的数据和遇到的错误.  

如果读取成功, 则 err 返回 nil, 而不是 EOF

`filename`必须是一个文件路径, 否则会出错

## 4. TempDir 

```go
TempDir func(dir, prefix string) (name string, err error)
```

操作系统中一般都会提供临时目录, 比如linux下的/tmp目录（通过os.TempDir()可以获取到). 

有时候, 我们自己需要创建临时目录, 比如Go工具链源码中（src/cmd/go/build.go）, 通过TempDir创建一个临时目录, 用于存放编译过程的临时文件.

## 5. TempFile 

```go
TempFile func(dir, prefix string) (f *os.File, err error)
```

在 dir 目录中创建一个以 prefix 为前缀的临时文件, 并将其以读写模式打开. 返回创建的文件对象和遇到的错误. 

如果 dir 为空, 则在默认的临时目录中创建文件（参见 os.TempDir）, 多次调用会创建不同的临时文件, 调用者可以通过 f.Name() 获取文件的完整路径.  

调用本函数所创建的临时文件, 应该由调用者自己删除. 

## 6. WriteFile 

```go
WriteFile func(filename string, data []byte, perm os.FileMode) error
```

向文件中写入数据, 写入前会**清空文件**.  

如果文件不存在, 则会以指定的权限创建该文件. 返回遇到的错误. 
# nodejs-执行命令行脚本(二)函数解析

参考文章

1. [在 Windows 上衍生 .bat 和 .cmd 文件](http://nodejs.cn/api/child_process.html#child_process_spawning_bat_and_cmd_files_on_windows)

2. [Node.js中spawn与exec的异同比较](https://segmentfault.com/a/1190000002913884)

`child_process`模块提供了衍生子进程的功能, 它与[popen(3)](http://man7.org/linux/man-pages/man3/popen.3.html) 类似, 但不完全相同. 这个功能主要由`child_process.spawn()` 函数提供. `child_process` 模块还提供了其他一些同步和异步的可选函数. **每个函数都是基于 `child_process.spawn()` 或 `child_process.spawnSync()` 实现的.**

- `spawn()`
- `exec()`
- `execFile()`
- `fork()`

每个函数都返回`ChildProcess`实例. 这些实例实现了Node.js `EventEmitter API`, 允许父进程注册监听器函数, 在子进程生命周期期间, 当特定的事件发生时会调用这些函数. `child_process.exec()` 和`child_process.execFile()` 函数可以额外指定一个可选的 callback 函数, 当子进程结束时会被调用.

`exec`和`spawn`在使用上只有传参格式, 回调绑定的方式有所区别(还更深层的不同可以见参考文章2).

```
spawn('bash', ['-c', 'npm', 'install']);
spawn('cmd.exe', ['/c', 'npm', 'install']);

exec('bash -c "npm install"');
exec('cmd.exe /c npm install');
```

`execFile()`倒是有所区别, 因为ta的原型为`child_process.execFile(file[, args][, options][, callback])`, ta的第一个参数为`file`, 就是目标脚本路径...目标脚本, 而不是目标命令, 不会在PATH中搜索中.

> nodejs官方文档上说: 在 Windows 上, `.bat` 和 `.cmd` 文件在没有终端的情况下是不可执行的, 因此不能使用 `child_process.execFile()` 启动.

但是!!!实验时是可以的, 见上一篇文档.

## 关于`fork()`

`child_process.fork()` 方法是 `child_process.spawn()` 的一个特殊情况, 专门用于衍生**新的Node.js进程**.

ta的函数原型为

```js
child_process.fork(modulePath[, args][, options])
```

最初我并不知道为什么其他函数的参数要么叫`file`, 要么叫`command`, 而`fork`的参数叫`modulePath`. 然后开启了一轮实验.

```js
var subproc = cprocess.fork('C:\\Program Files\\MxiPlayer-3.0.8\\MxiPlayer.exe', ['config.json'], {cwd: 'C:\\Program Files\\MxiPlayer-3.0.8'});
console.log(subproc);
```

执行失败, 瞄了一眼打印的日志, 正好看到如下的内容.

```
  spawnargs:
   [ 'C:\\Program Files\\nodejs\\node.exe',
     'C:\\Program Files\\MxiPlayer-3.0.8\\MxiPlayer.exe',
     'config.json' ],
```

好像是用`node`来执行的??? node调用exe, bat肯定不行的啊.

然后回头再看官方文档...嗯, `cprocess.fork('./test.js');`这样用才对.

其实fork()也可以执行shell/bat脚本, 方法见上一篇文档.
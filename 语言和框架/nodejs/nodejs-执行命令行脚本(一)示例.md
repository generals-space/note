# nodejs-执行命令行脚本(一)示例

场景描述

win平台下, 有一个bat脚本`MxiPlayer.exe`(没有办法查看ta的源码, 但ta的运行方式和运行表现都和普通的bat脚本没有区别, 就把ta当作普通脚本来对待), 启动时需要指定配置文件, 启动后会打开cmd命令行, 打印一些日志, 执行完毕后cmd关闭.

```
./MxiPlayer.exe ./config.json
```

由于我把这个脚本放在了`C:\Program Files`目录下, 路径中有空格, 需要用引号包裹起来, 参数也是, 所以增加了命令和程序的书写难度. 

在cmd和powershell下, 可用的命令各自如下

```bat
; 注意: 必须为双引号, 单引号会出错
"C:\Program Files\MxiPlayer-3.0.8\MxiPlayer.exe" "C:\Program Files\MxiPlayer-3.0.8\config.json"

```ps1
// 前面的&符号不能删掉
& 'C:\Program Files\MxiPlayer-3.0.8\MxiPlayer.exe' 'C:\Program Files\MxiPlayer-3.0.8\config.json'
```

------

在nodejs中, 使用`child_process`标准库可以调用命令行脚本. 下面的示例中我们使用了4种函数, 都是经过测试的.

## 1. `exec()`

这是唯一一个bat脚本执行完毕后可以自动结束的函数, 其他的都不行.

```js
// 以下两个command都可以
const command = '"C:\\Program Files\\MxiPlayer-3.0.8\\MxiPlayer.exe" "C:\\Program Files\\MxiPlayer-3.0.8\\config.json"';
// const command = 'MxiPlayer.exe config.json';
cprocess.exec(command, {cwd: 'C:\\Program Files\\MxiPlayer-3.0.8'}, (err, stdout, stderr) => {
    if(err){
        console.error(err);
        return;
    }
    console.log(stdout);
});
```

## 2. `spawn()`

```js
// 以下两种都可以, 由于目标程序其实也是脚本, 所以不指定`cmd.exe`也可以成功执行.
const bat = cprocess.spawn('cmd.exe', ['/c', 'MxiPlayer.exe', 'config.json'], {cwd: 'C:\\Program Files\\MxiPlayer-3.0.8'});
// const bat = cprocess.spawn('MxiPlayer.exe', ['config.json'], {cwd: 'C:\\Program Files\\MxiPlayer-3.0.8'});

bat.stdout.on('data', (data) => {console.log(data.toString());});
bat.stderr.on('data', (data) => {console.log(data.toString());});
bat.on('exit', (code) => {console.log(`子进程退出码：${code}`);});
```

## 3. `execFile()`

nodejs官方文档说在win平台下无法使用这个函数, 因为无法创建子shell来执行脚本, 但是经过实验是可以的.

```js
const dir = 'C:\\Program Files\\MxiPlayer-3.0.8';
const command = 'C:\\Program Files\\MxiPlayer-3.0.8\\MxiPlayer.exe';
const config = 'C:\\Program Files\\MxiPlayer-3.0.8\\config.json';
// 以下两种execFile都可以
// const child = cprocess.execFile('MxiPlayer.exe', ['./config.json'], {cwd: dir}, (error, stdout, stderr) => {
const child = cprocess.execFile(command, [config], {cwd: dir}, (error, stdout, stderr) => {
    if (error) {
        throw error;
    }
    console.log(stdout);
});
```

## 4. `fork()`

nodejs里的fork和其他语言中的fork不一样, 不需要在父子两个进程间分别处理逻辑. ta只是`spawn`的一个特殊情况, ta只是用来开一个子进程执行一个`js`脚本...但是在我知道这一点前我已经先把脚本运行起来了...

```js
// 以下两种fork都可以
// 第一个参数随便了, 因为执行cmd脚本时不能使用第一个参数来指定, 而是需要使用选项参数中的`execPath`
var subproc = cprocess.fork('xxx', ['config.json'], {cwd: 'C:\\Program Files\\MxiPlayer-3.0.8', execPath: 'MxiPlayer.exe'});
// var subproc = cprocess.fork('xxx', ['config.json'], {cwd: 'C:\\Program Files\\MxiPlayer-3.0.8', execPath: 'C:\\Program Files\\MxiPlayer-3.0.8\\MxiPlayer.exe'});

console.log(subproc);
```
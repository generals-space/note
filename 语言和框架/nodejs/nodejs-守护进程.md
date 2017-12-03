# NodeJS-守护进程

参考文章

1. [nodejs如何以守护进程运行啊？](http://cnodejs.org/topic/4f1d05b4817ae4105c033e38)

2. [Nodejs编写守护进程](https://ashan.org/archives/917)

3. [nodejs写自己的守护进程 防止进程死掉](http://cnodejs.org/topic/5218d6b5bee8d3cb12540653)

```js
var cluster = require('cluster');
if (cluster.isMaster) {
  //Fork a worker to run the main program
  for (var i = 0; i < 2; i++) var worker = cluster.fork();
} else {
  //Run main program
  require('./app.js');
  console.log('worker is running');
}

cluster.on('death', function(worker) {
  //If the worker died, fork a new worker
  console.log('worker ' + worker.pid + ' died. restart...');
  cluster.fork();
});
```
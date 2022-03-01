```
java -XX:+ -version
```

## jinfo -flags $pid

可以查看目标进程的 jvm 启动参数(其实使用`ps -ef | grep java`也能得到这样的结果吧)

```console
$ jinfo -flags 13
Attaching to process ID 13, please wait...
Debugger attached successfully.
Server compiler detected.
JVM version is 25.232-b09

Non-default VM flags: -XX:CICompilerCount=2 -XX:CMSInitiatingOccupancyFraction=75 -XX:+HeapDumpOnOutOfMemoryError -XX:InitialHeapSize=1073741824 -XX:MaxHeapSize=1073741824 -XX:MaxNewSize=268435456 -XX:MaxTenuringThreshold=6 -XX:MinHeapDeltaBytes=196608 -XX:NewSize=268435456 -XX:OldPLABSize=16 -XX:OldSize=805306368 -XX:+UseCMSInitiatingOccupancyOnly -XX:+UseCompressedClassPointers -XX:+UseCompressedOops -XX:+UseConcMarkSweepGC -XX:+UseParNewGC

Command line: -XX:+UseConcMarkSweepGC -XX:CMSInitiatingOccupancyFraction=75 -XX:+UseCMSInitiatingOccupancyOnly -Djava.awt.headless=true -Dfile.encoding=UTF-8 -Djruby.compile.invokedynamic=true -Djruby.jit.threshold=0 -Djruby.regexp.interruptible=true -XX:+HeapDumpOnOutOfMemoryError -Djava.security.egd=file:/dev/urandom -Dlog4j2.isThreadContextMapInheritable=true -XX:NewSize=256m -Xmx1024m -Xms1024m
```

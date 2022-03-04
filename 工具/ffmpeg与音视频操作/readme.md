#ffmpeg与音视频操作

参考文章

1. [ffprobe,ffplay ffmpeg常用的命令行命令](https://juejin.im/post/5a59993cf265da3e4f0a1e4b)

ffmpeg不像protobuf工具或是redis等中间件一样拥有各种语言的SDK, 网络上大部分操作都是通过`exec`函数族通过命令调用的, 并且执行结果也都是从标准输出中读取然后用正则等手段截取的. 所以这个系统中主要是讲ffmpeg相关命令的使用方法.

ffmpeg工具集包含如下常用工具命令

1. ffprobe 用于查看媒体文件信息;

2. ffplay 用于播放音视频;

3. ffmpeg 用于音视频的剪辑, 拼接, 转换等操作;
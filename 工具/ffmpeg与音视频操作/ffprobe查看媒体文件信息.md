# ffprobe查看媒体文件信息

参考文章

1. [ffprobe,ffplay ffmpeg常用的命令行命令](https://juejin.im/post/5a59993cf265da3e4f0a1e4b)

2. [python获取音频的长度ffmpeg或ffprobe](https://wqian.net/blog/2018/1128-python-ffmpeg-mp3-length-index.html)

## 1. 查看文件信息

首先查看一个音频文件

```
$ ffprobe 永夜\ -\ 石岩.mp3
ffprobe version 4.1 Copyright (c) 2007-2018 the FFmpeg developers
...省略
[mp3 @ 0x7ff7a5800000] Estimating duration from bitrate, this may be inaccurate
Input #0, mp3, from '永夜 - 石岩.mp3':
  Metadata:
    title           : yongye1_2
    date            : 2018
    track           : 1
  Duration: 00:04:12.44, start: 0.000000, bitrate: 320 kb/s
    Stream #0:0: Audio: mp3, 48000 Hz, stereo, fltp, 320 kb/s
```

`Duration`, `start`, `bitrate`一行分别表示音频的**时长**, **开始播放时间**和**比特率**.

`Stream`, `Audio`一行可以看出, 该音频流**编码格式**是mp3, **采样率**是44.1khz, **采样表示格式**(不懂...), 这路流的比特率是320kb/s.

------

然后是一个视频文件

```
$ ffprobe RWBY\ White\ trailer\ -\ Miku\ version.mp4
ffprobe version 4.1 Copyright (c) 2007-2018 the FFmpeg developers
...省略
Input #0, mov,mp4,m4a,3gp,3g2,mj2, from 'RWBY White trailer - Miku version.mp4':
  Metadata:
    major_brand     : isom
    minor_version   : 1
    compatible_brands: isom
    creation_time   : 2013-08-09T12:20:13.000000Z
    copyright       :
    copyright-eng   :
  Duration: 00:03:47.04, start: 0.000000, bitrate: 1406 kb/s
    Stream #0:0(und): Video: h264 (High) (avc1 / 0x31637661), yuv420p(tv, bt709/unknown/unknown), 1280x720, 1282 kb/s, 60 fps, 60 tbr, 5000k tbn, 120 tbc (default)
    Metadata:
      creation_time   : 2013-08-08T16:18:01.000000Z
      handler_name    : Video
    Stream #0:1(und): Audio: aac (LC) (mp4a / 0x6134706D), 44100 Hz, stereo, fltp, 117 kb/s (default)
    Metadata:
      creation_time   : 2013-08-08T17:03:03.000000Z
      handler_name    : Audio
```

`Stream`, `Video`一行表示, 视频流编码方式为`h264`, 每一帧都是`yuv420p`格式, 分辨率是`1280x720`, 每秒60帧.

`Stream`, `Audio`一行表示, 音频流编码方式为`aac`, (封装格式是LC(不懂...)), 采样率是44.1khz.

## 2. 输出内容

上面通过`ffprobe`直接查询的媒体文件太过杂乱, `ffprobe`提供了一些参数可以输出格式化的信息.

### 2.1 `-show_format`: 输出媒体文件的格式信息(包括时长, 文件大小, 格式信息等等)

```
$ ffprobe -show_format 永夜\ -\ 石岩.mp3
ffprobe version 4.1 Copyright (c) 2007-2018 the FFmpeg developers
...省略
[mp3 @ 0x7fd6ba000000] Estimating duration from bitrate, this may be inaccurate
Input #0, mp3, from '永夜 - 石岩.mp3':
  Metadata:
    title           : yongye1_2
    date            : 2018
    track           : 1
  Duration: 00:04:12.44, start: 0.000000, bitrate: 320 kb/s
    Stream #0:0: Audio: mp3, 48000 Hz, stereo, fltp, 320 kb/s
[FORMAT]
filename=永夜 - 石岩.mp3
nb_streams=1
nb_programs=0
format_name=mp3
format_long_name=MP2/3 (MPEG audio layer 2/3)
start_time=0.000000
duration=252.435200
size=10097408
bit_rate=320000
probe_score=51
TAG:title=yongye1_2
TAG:date=2018
TAG:track=1
[/FORMAT]
```

### 2.2 `-show_streams`: 输出每个流最详细的信息,例如视频宽高信息,是否有B帧,视频帧的总数目,编码格式,显示比例,音频的省道,编码格式等等.

```
$ ffprobe -show_streams 永夜\ -\ 石岩.mp3
...这里省略
```

### 2.3 `-show_frames`: 显示帧信息(不懂...)

```
$ ffprobe -show_frames 永夜\ -\ 石岩.mp3
```

### 2.4 `-show_packets`: 查看包信息(不懂...)

```
$ ffprobe -show_packets 永夜\ -\ 石岩.mp3
```

## 3. 输出格式

上述输出格式相同, 类似于`ini`文件格式, 其实还有更多标准的格式, 查看`ffprobe --help`可以发现如下选项.

```
-print_format format  set the output printing format (available formats are: default, compact, csv, flat, ini, json, xml)
```

可以看到`-print_format`选项可以指定`int`, `json`, `xml`等多种类型的格式.

示例如下

```
$ ffprobe -show_format -print_format json 永夜\ -\ 石岩.mp3
ffprobe version 4.1 Copyright (c) 2007-2018 the FFmpeg developers
...省略
{
[mp3 @ 0x7fd41b004c00] Estimating duration from bitrate, this may be inaccurate
Input #0, mp3, from '永夜 - 石岩.mp3':
  Metadata:
    title           : yongye1_2
    date            : 2018
    track           : 1
  Duration: 00:04:12.44, start: 0.000000, bitrate: 320 kb/s
    Stream #0:0: Audio: mp3, 48000 Hz, stereo, fltp, 320 kb/s
    "format": {
        "filename": "永夜 - 石岩.mp3",
        "nb_streams": 1,
        "nb_programs": 0,
        "format_name": "mp3",
        "format_long_name": "MP2/3 (MPEG audio layer 2/3)",
        "start_time": "0.000000",
        "duration": "252.435200",
        "size": "10097408",
        "bit_rate": "320000",
        "probe_score": 51,
        "tags": {
            "title": "yongye1_2",
            "date": "2018",
            "track": "1"
        }
    }
}
```

但是还有大段的无用信息, 这在使用程序中exec调用时对标准输出的解析是非常不友好的. 那么怎么去掉呢?

`-v`和`-loglevel`选项, 这两个选项完全一样, 不知道有多少可用参数, 但是至少有一个是确定的: `quite`.

```
$ ffprobe -show_format -print_format json -v quiet 永夜\ -\ 石岩.mp3
{
    "format": {
        "filename": "永夜 - 石岩.mp3",
        "nb_streams": 1,
        "nb_programs": 0,
        "format_name": "mp3",
        "format_long_name": "MP2/3 (MPEG audio layer 2/3)",
        "start_time": "0.000000",
        "duration": "252.435200",
        "size": "10097408",
        "bit_rate": "320000",
        "probe_score": 51,
        "tags": {
            "title": "yongye1_2",
            "date": "2018",
            "track": "1"
        }
    }
}
```

这样在读取的时候就容易多了.

# ffmpeg分离mp4视频中的mp3音频

参考文章

1. [采用FFmpeg从视频中提取音频(声音)保存为mp3文件](https://blog.csdn.net/gobitan/article/details/50489339)

参考文章1中说ffmpeg只有mp3解码, 没有mp3编码, 所以想转出mp3还需要编译额外的编码库. 

实际上是不需要的, 评论中有人指出了.

实验机器: CentOS 7
ffmpeg版本: 2.8.15

从mp4视频中分离mp3音频

```
ffmpeg -i xiaonan.mp4 -f mp3 -vn xiaonan.mp3
ffmpeg version 2.8.15 Copyright (c) 2000-2018 the FFmpeg developers
  built with gcc 4.8.5 (GCC) 20150623 (Red Hat 4.8.5-36)
  configuration: --prefix=/usr --bindir=/usr/bin --datadir=/usr/share/ffmpeg --incdir=/usr/include/ffmpeg --libdir=/usr/lib64 --mandir=/usr/share/man --arch=x86_64 --optflags='-O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector-strong --param=ssp-buffer-size=4 -grecord-gcc-switches -m64 -mtune=generic' --extra-ldflags='-Wl,-z,relro ' --enable-libopencore-amrnb --enable-libopencore-amrwb --enable-libvo-amrwbenc --enable-version3 --enable-bzlib --disable-crystalhd --enable-gnutls --enable-ladspa --enable-libass --enable-libcdio --enable-libdc1394 --enable-libfdk-aac --enable-nonfree --disable-indev=jack --enable-libfreetype --enable-libgsm --enable-libmp3lame --enable-openal --enable-libopenjpeg --enable-libopus --enable-libpulse --enable-libschroedinger --enable-libsoxr --enable-libspeex --enable-libtheora --enable-libvorbis --enable-libv4l2 --enable-libx264 --enable-libx265 --enable-libxvid --enable-x11grab --enable-avfilter --enable-avresample --enable-postproc --enable-pthreads --disable-static --enable-shared --enable-gpl --disable-debug --disable-stripping --shlibdir=/usr/lib64 --enable-runtime-cpudetect
  libavutil      54. 31.100 / 54. 31.100
  libavcodec     56. 60.100 / 56. 60.100
  libavformat    56. 40.101 / 56. 40.101
  libavdevice    56.  4.100 / 56.  4.100
  libavfilter     5. 40.101 /  5. 40.101
  libavresample   2.  1.  0 /  2.  1.  0
  libswscale      3.  1.101 /  3.  1.101
  libswresample   1.  2.101 /  1.  2.101
  libpostproc    53.  3.100 / 53.  3.100
Input #0, mov,mp4,m4a,3gp,3g2,mj2, from 'xiaonan.mp4':
  Metadata:
    major_brand     : mp42
    minor_version   : 0
    compatible_brands: isommp42
    creation_time   : 2017-05-25 15:22:05
  Duration: 00:02:30.70, start: 0.000000, bitrate: 691 kb/s
    Stream #0:0(und): Video: h264 (Constrained Baseline) (avc1 / 0x31637661), yuv420p(tv, bt709), 640x360 [SAR 1:1 DAR 16:9], 593 kb/s, 25 fps, 25 tbr, 12800 tbn, 50 tbc (default)
    Metadata:
      handler_name    : VideoHandler
    Stream #0:1(und): Audio: aac (LC) (mp4a / 0x6134706D), 44100 Hz, stereo, fltp, 96 kb/s (default)
    Metadata:
      creation_time   : 2017-05-25 15:22:06
      handler_name    : IsoMedia File Produced by Google, 5-11-2011
Output #0, mp3, to 'xiaonan.mp3':
  Metadata:
    major_brand     : mp42
    minor_version   : 0
    compatible_brands: isommp42
    TSSE            : Lavf56.40.101
    Stream #0:0(und): Audio: mp3 (libmp3lame), 44100 Hz, stereo, fltp (default)
    Metadata:
      creation_time   : 2017-05-25 15:22:06
      handler_name    : IsoMedia File Produced by Google, 5-11-2011
      encoder         : Lavc56.60.100 libmp3lame
Stream mapping:
  Stream #0:1 -> #0:0 (aac (native) -> mp3 (libmp3lame))
Press [q] to stop, [?] for help
size=    2355kB time=00:02:30.70 bitrate= 128.0kbits/s
video:0kB audio:2355kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: 0.014347%
```

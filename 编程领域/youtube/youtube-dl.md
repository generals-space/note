## 最简示例

```
youtube-dl --proxy socks5://127.0.0.1:1080 --no-check-certificate --username huangjiaruo321@gmail.com --password 123456 https://www.youtube.com/watch?v=ByCDKBBhdZg
```

定义输出格式, 清晰度, 文件名

```
youtube-dl --proxy socks5://127.0.0.1:1080 --no-check-certificate --username huangjiaruo321@gmail.com --password 123456 --format best -o '%(upload_date)s-%(title)s.%(ext)s' --write-auto-sub https://www.youtube.com/watch?v=ByCDKBBhdZg
```

## 列表下载

```
youtube-dl --proxy socks5://127.0.0.1:1080 --no-check-certificate --username huangjiaruo321@gmail.com --password 123456 --format best -o '%(playlist_index)05d-%(upload_date)s-%(title)s.%(ext)s' --write-auto-sub https://www.youtube.com/playlist?list=PLsRn8zzjiZRjpuTNLrzCFtwgbRe8SyNB5
```

由于失败导致的中断可以通过`--playlist-start 10`指定继续下载的起始索引.

`--retries 数字/infinite`设置重试次数...好像不管用?

```
$ youtube-dl --proxy socks5://192.168.0.5:1080 --no-check-certificate --username huangjiaruo321@gmail.com --password 123456 --playlist-start 22 --retries infinite --format best -o '%(upload_date)s-%(title)s.%(ext)s' --write-auto-sub https://www.youtube.com/playlist?list=PLsRn8zzjiZRjpuTNLrzCFtwgbRe8SyNB5
[download] Downloading video 10 of 276
[youtube] 4vY0NqKNmB4: Downloading webpage
[youtube] 4vY0NqKNmB4: Downloading video info webpage
[youtube] 4vY0NqKNmB4: Looking for automatic captions
ERROR: 失败...
```

关于失败后, 下次起始索引的计算, 其值应该为上次`--playlist-start`值, 加上本次失败的索引, 再减1.

比如上面从索引22开始, 下载到第10个视频时失败, 那么下次开始时应该指定索引为`22 + 10 - 1 = 31`.

## 关于格式与清晰度

查看目标视频可选的格式, 包括清晰度, 文件格式等, 不会下载.

```
$ youtube-dl --proxy socks5://127.0.0.1:1080 --no-check-certificate --username huangjiaruo321@gmail.com --password 123456 --list-formats https://www.youtube.com/watch?v=ByCDKBBhdZg
[youtube] Downloading login page
[youtube] Looking up account info
WARNING: Unable to look up account info: <urlopen error EOF occurred in violation of protocol (_ssl.c:1051)>
[youtube] ByCDKBBhdZg: Downloading webpage
[youtube] ByCDKBBhdZg: Downloading video info webpage
[info] Available formats for ByCDKBBhdZg:
format code  extension  resolution note
250          webm       audio only tiny   71k , opus @ 70k (48000Hz), 2.30MiB
140          m4a        audio only tiny  130k , m4a_dash container, mp4a.40.2@128k (44100Hz), 5.18MiB
...省略
22           mp4        1280x720   720p  551k , avc1.64001F, mp4a.40.2@192k (44100Hz) (best)
```

## 输出文件名(OUTPUT TEMPLATE)

默认存储在本地的文件名为, 视频title + 视频id(url中的那个).后缀(一般是mp4)

`-o '%(upload_date)s-%(title)s.%(ext)s'` 注意引号

## 关于字幕

`--list-subs [url]` : 列出所有可用字幕
`--write-sub [url]`: 这样会下载一个vtt格式的英文字幕和mkv格式的1080p视频下来, 但是如果没有上传的字幕, 就不会下载.
`--write-auto-sub`: 只对youtube视频有效.
`--write-sub --skip-download [url]`: 下载单独的vtt字幕文件,而不会下载视频

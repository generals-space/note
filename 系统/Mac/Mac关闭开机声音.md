# Mac关闭开机声音

参考文章

1. [如何消除苹果Mac电脑开机声音](https://www.jianshu.com/p/14a29719dcda)

好像没用.

`sound_off.sh`

```bash
#!/bin/bash
osascript -e 'set volume output muted 1'
```

`sound_on.sh`

```bash
#!/bin/bash
osascript -e 'set volume output muted 0'
```

把`sound_off.sh`和`sound_on.sh`这两个脚本放在`/usr/local/bin/`和`/Library/Scripts/`都不行...

```
sudo defaults write com.apple.loginwindow LogoutHook /Library/Scripts/sound_off.sh
sudo defaults write com.apple.loginwindow LoginHook /Library/Scripts/sound_on.sh
```

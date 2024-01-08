# MacOS m1系统code命令执行失败(转)

参考文章

1. [MacOS 12.3 无法正常使用code命令的解决方法](https://www.wyr.me/post/692)
    - 原文链接
2. [fix: remove python usage in macOS cli wrapper](https://github.com/microsoft/vscode/pull/138582)

MacOS: 12.6.3(m1 pro)
VSCode: 版本: 1.64.2

~~从MacOS 12.3 Beta版本开始，系统将不再内置python2且将无法正常安装python2，无论是intel芯片还是Apple芯片的设备都无法安装。原因是/usr/bin/python的软链接无法正常被删除或覆盖。并且默认不开启python3命令。~~

**2022年04月17日14:58:00更新：** 从MacOS 12.4 Beta版(21F5048e) 开始，可以通过pyenv在intel和Apple芯片中安装python2。详细方法见《brew安装python2》。

因此可能会导致一系列依赖python命令的应用程序无法运行。

例如将会遇到VS(Visual Studio Code)无法使用code命令。

```console
$ code .
/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code: line 6: python: command not found
/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code: line 10: ./MacOS/Electron: No such file or directory
```

由此也可能导致"Visual Studio Code - Insiders" needs to be updated on macOS Monterey弹窗的问题。

相关问题及VS Code最终Merge的方案见链接：https://github.com/microsoft/vscode/pull/138582

即, 使用纯Shell方式代替python命令作为优选的方案。

在VS Code正式更新此代码之前可以通过修改/usr/local/bin/code文件解决：

```bash
#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.

function app_realpath() {
    SOURCE=$1
    while [ -h "$SOURCE" ]; do
        DIR=$(dirname "$SOURCE")
        SOURCE=$(readlink "$SOURCE")
        [[ $SOURCE != /* ]] && SOURCE=$DIR/$SOURCE
    done
    SOURCE_DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
    echo "${SOURCE_DIR%%${SOURCE_DIR#*.app}}"
}

APP_PATH="$(app_realpath "${BASH_SOURCE[0]}")"
if [ -z "$APP_PATH" ]; then
    echo "Unable to determine app path from symlink : ${BASH_SOURCE[0]}"
    exit 1
fi
CONTENTS="$APP_PATH/Contents"
ELECTRON="$CONTENTS/MacOS/Electron"
CLI="$CONTENTS/Resources/app/out/cli.js"
ELECTRON_RUN_AS_NODE=1 "$ELECTRON" "$CLI" --ms-enable-electron-run-as-node "$@"
exit $?
```

------

补充一下原本的`code`命令内容

```bash
#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.

function realpath() { python -c "import os,sys; print(os.path.realpath(sys.argv[1]))" "$0"; }
CONTENTS="$(dirname "$(dirname "$(dirname "$(dirname "$(realpath "$0")")")")")"
ELECTRON="$CONTENTS/MacOS/Electron"
CLI="$CONTENTS/Resources/app/out/cli.js"
ELECTRON_RUN_AS_NODE=1 "$ELECTRON" "$CLI" --ms-enable-electron-run-as-node "$@"
exit $?
```

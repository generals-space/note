# sublime插件

## 1. sublime插件安装/卸载

### 1.1 安装方法

1.按`Ctrl+\``调出console

2.粘贴以下代码到底部命令行并回车：

```py
import urllib.request,os; pf = 'Package Control.sublime-package'; ipp = sublime.installed_packages_path(); urllib.request.install_opener( urllib.request.build_opener( urllib.request.ProxyHandler()) ); open(os.path.join(ipp, pf), 'wb').write(urllib.request.urlopen( 'http://sublime.wbond.net/' + pf.replace(' ','%20')).read())
```

3.重启Sublime Text 3， 如果在`Perferences->package settings`中看到`package control`这一项，则安装成功。

以下是`sublime text2`的安装代码：

```py
import urllib2,os; pf='Package Control.sublime-package'; ipp = sublime.installed_packages_path(); os.makedirs( ipp ) if not os.path.exists(ipp) else None; urllib2.install_opener( urllib2.build_opener( urllib2.ProxyHandler( ))); open( os.path.join( ipp, pf), 'wb' ).write( urllib2.urlopen( 'http://sublime.wbond.net/' +pf.replace( ' ','%20' )).read()); print( 'Please restart Sublime Text to finish installation')
```

4.`Ctrl+Shift+p`调用`Package Control`工具，输入`Install Package`，回车，然后输入要安装的插件的名称即可。

### 1.2 卸载插件

`Ctrl+Shift+p`，输入`remove package`，然后输入要卸载的插件名称即可.

## 2. 插件推荐列表

(1)BracketHighlighter

类似于代码匹配，可以匹配括号，引号等符号内的范围。

(2)SublimeGit

Git插件，可适用于sublime text3

(3)

markdown插件, 语法高亮与预览

markdown editing, markdown preview, markdown extended, monokai extended.

(4)CoolFormat

简单好用的代码格式化工具，相当于简化版的Astyle，默认ctrl+alt+shift+q格式化当前文件，ctrl+alt+shift+s格式化当前选中。

## 3. markdown插件

### 3.1. 说明

1. 首先安装markdown插件时最好不要正在编辑md文件

2. sublime的markdown插件也无法实现实时预览, 只能保存, 然后刷新浏览器

### 3.2. 所需的插件

- markdown editing

- markdown preview

- markdown extended

- monokai extended.


### 3.3 主题设置

未设置主题之前会出现很奇怪的情况, 其他标签页显示正常, md文件将会呈现一种惨白的颜色, 十分晃眼;

Preference->Color Scheme->Monokai Extended->Monokai Extended;

### 3.4 高亮语法选择

View->Syntax->Open all with current extension as...-> markdown extended

### 3.5 预览/生成html文件

Ctrl + Shift + p, 输入markdown会有提示, 选择Preview in browser / Save in html

## 4. Bracket Highlighter

BracketHighlighter配置

Pregerence > Package Settings > BracketHighlighter > Bracket Settings – User

对应的关系：

- {} － curly

- () － round

- [] － square

- <> － angle

- “” ” － quote

style: solid、underline、outline、highlight

```json
{
    "bracket_styles": {
        "default": {
            "icon": "dot",
            // "color": "entity.name.class",
            "color": "brackethighlighter.default",
            "style": "underline"
        },
 
        "unmatched": {
            "icon": "question",
            "color": "brackethighlighter.unmatched",
            "style": "underline"
        },
        "curly": {
            "icon": "curly_bracket",
            "color": "brackethighlighter.curly",
            "style": "underline"
        },
        "round": {
            "icon": "round_bracket",
            //"color": "brackethighlighter.round",
            "color": "brackethighlighter.quote",
            "style": "underline"
        },
        "square": {
            "icon": "square_bracket",
            "color": "brackethighlighter.square",
            "style": "underline"
        },
        "angle": {
            "icon": "angle_bracket",
            "color": "brackethighlighter.angle",
            "style": "underline"
        },
        "tag": {
            "icon": "tag",
            "color": "brackethighlighter.tag",
            "style": "underline"
        },
        "single_quote": {
            "icon": "single_quote",
            "color": "brackethighlighter.quote",
            "style": "underline"
        },
        "double_quote": {
            "icon": "double_quote",
            "color": "brackethighlighter.quote",
            "style": "underline"
        },
        "regex": {
            "icon": "regex",
            "color": "brackethighlighter.quote",
            "style": "outline"
        }
    }
 
}
```

### 修改配色方案

参考文章：

[Sublime Text3 BracketHighlighter色彩配置](http://www.dbpoo.com/sublime-text3-brackethighlighter/)

[SublimeText插件BracketHighlighter配置](http://www.darkpool.net/archives/95)

找到Sublime text3安装目录下的Packages中的`Color Scheme – Default.sublime-package`

Linux下`/opt/sublime_text_3/Packages/Color Scheme - Default.sublime-package`

添加后缀名`Color Scheme – Default.sublime-package.zip`，解压，找到`Monokai.tmTheme`，修改。

```xml
<!-- Bracket 开始 -->
<dict>
    <key>name</key>
    <string>Bracket Default</string>
    <key>scope</key>
    <string>brackethighlighter.default</string>
    <key>settings</key>
    <dict>
        <key>foreground</key>
        <string>#FFFFFF</string>
        <key>background</key>
        <string>#A6E22E</string>
    </dict>
</dict>
 
<dict>
    <key>name</key>
    <string>Bracket Unmatched</string>
    <key>scope</key>
    <string>brackethighlighter.unmatched</string>
    <key>settings</key>
    <dict>
        <key>foreground</key>
        <string>#FFFFFF</string>
        <key>background</key>
        <string>#FF0000</string>
    </dict>
</dict>
 
<dict>
    <key>name</key>
    <string>Bracket Curly</string>
    <key>scope</key>
    <string>brackethighlighter.curly</string>
    <key>settings</key>
    <dict>
        <key>foreground</key>
        <string>#FF00FF</string>
    </dict>
</dict>
 
<dict>
    <key>name</key>
    <string>Bracket Round</string>
    <key>scope</key>
    <string>brackethighlighter.round</string>
    <key>settings</key>
    <dict>
        <key>foreground</key>
        <string>#E7FF04</string>
    </dict>
</dict>
 
<dict>
    <key>name</key>
    <string>Bracket Square</string>
    <key>scope</key>
    <string>brackethighlighter.square</string>
    <key>settings</key>
    <dict>
        <key>foreground</key>
        <string>#FE4800</string>
    </dict>
</dict>
 
<dict>
    <key>name</key>
    <string>Bracket Angle</string>
    <key>scope</key>
    <string>brackethighlighter.angle</string>
    <key>settings</key>
    <dict>
        <key>foreground</key>
        <string>#02F78E</string>
    </dict>
</dict>
 
<dict>
    <key>name</key>
    <string>Bracket Tag</string>
    <key>scope</key>
    <string>brackethighlighter.tag</string>
    <key>settings</key>
    <dict>
        <key>foreground</key>
        <string>#FFFFFF</string>
        <key>background</key>
        <string>#0080FF</string>
    </dict>
</dict>
 
<dict>
    <key>name</key>
    <string>Bracket Quote</string>
    <key>scope</key>
    <string>brackethighlighter.quote</string>
    <key>settings</key>
    <dict>
        <key>foreground</key>
        <string>#56FF00</string>
    </dict>
</dict>
<!-- Bracket 结束 -->
```

将上边的代码添加到`Monokai.tmTheme`中，注意添加的位置，与其中大部分`<dict></dict>`并列就可以。
然后再将修改完成的文件放到压缩包`Color Scheme – Default.sublime-package.zip`里边，改名 `Color Scheme – Default.sublime-package`放回`/opt/sublime_text3/Packages`中。
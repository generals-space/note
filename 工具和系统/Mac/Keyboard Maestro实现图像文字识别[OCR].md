# Keyboard Maestro实现图像文字识别[OCR]

参考文章

1. [Keyboard Maestro | 使用 KM 实现免费中文 OCR 光学字符识别](https://zhuanlan.zhihu.com/p/60286600)
	- MacOS上的4个截屏快捷键
2. [在 Mac 上随时提取屏幕上的文字和解析二维码](https://zhuanlan.zhihu.com/p/143162320)
	- 文字识别是根据 Keyboard Maestro 内置的
    - 文字识别我用的是参考文章1的脚本, 这个只看了二维码识别
3. [OCR Python SDK](https://ai.baidu.com/ai-doc/OCR/wkibizyjk)
    - 百度 ocr 接口的 python sdk 官方文档

Keyboard Maestro: v9.0.6

安装`Keyboard Maestro`后, 将如下代码保存为后缀名为`.kmmacros`的文件, 双击即可导入.

快捷键是`Ctrl+Command+Shift+A`, 选中屏幕上任意区域后, 会将该区域中的文字识别出来并放到系统剪切板中, 直接粘贴就可以了.

需要注意的是, 图像识别功能借助了百度OCR服务, 个人经过实名认证后, 拥有1000次/月的API调用额度, 超出后会收费.

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<array>
	<dict>
		<key>Activate</key>
		<string>Normal</string>
		<key>CreationDate</key>
		<real>575169677.49306595</real>
		<key>Macros</key>
		<array>
			<dict>
				<key>Actions</key>
				<array>
					<dict>
						<key>ActionName</key>
						<string>触发截图，自行根据截图快捷键修改</string>
						<key>KeyCode</key>
						<integer>21</integer>
						<key>MacroActionType</key>
						<string>SimulateKeystroke</string>
						<key>Modifiers</key>
						<integer>4864</integer>
						<key>ReleaseAll</key>
						<false/>
						<key>TargetApplication</key>
						<dict/>
						<key>TargetingType</key>
						<string>Front</string>
					</dict>
					<dict>
						<key>ActionName</key>
						<string>直到剪贴板内容为图片时继续</string>
						<key>ActionNotes</key>
						<string>测试</string>
						<key>Conditions</key>
						<dict>
							<key>ConditionList</key>
							<array>
								<dict>
									<key>ClipboardConditionType</key>
									<string>HasImage</string>
									<key>ClipboardText</key>
									<string></string>
									<key>ConditionType</key>
									<string>Clipboard</string>
								</dict>
							</array>
							<key>ConditionListMatch</key>
							<string>All</string>
						</dict>
						<key>MacroActionType</key>
						<string>PauseUntil</string>
						<key>TimeOutAbortsMacro</key>
						<true/>
					</dict>
					<dict>
						<key>ActionName</key>
						<string>保存截图至本地文件夹</string>
						<key>Append</key>
						<false/>
						<key>Destination</key>
						<string>~/Pictures/ocr-img.png</string>
						<key>Encoding</key>
						<string>UTF8</string>
						<key>Format</key>
						<string>PNG</string>
						<key>Format2</key>
						<string>PNG</string>
						<key>MacroActionType</key>
						<string>WriteFile</string>
						<key>Source</key>
						<string>Clipboard</string>
					</dict>
					<dict>
						<key>ActionName</key>
						<string>执行代码</string>
						<key>DisplayKind</key>
						<string>Clipboard</string>
						<key>HonourFailureSettings</key>
						<true/>
						<key>IncludeStdErr</key>
						<false/>
						<key>MacroActionType</key>
						<string>ExecuteShellScript</string>
						<key>Path</key>
						<string></string>
						<key>Source</key>
						<string>Nothing</string>
						<key>Text</key>
						<string>#!/usr/local/bin/python3

## pip install baidu-aip=2.2.18.0
from aip import AipOcr
import sys,io
sys.stdout=io.TextIOWrapper(sys.stdout.buffer,encoding='utf8')

## 需要先在控制台创建应用, 然后才能得到如下3个参数
APP_ID      = '24796865'
API_KEY     = 'h7Xbb7eRZGSruRS4ZfcPKb33'
SECRET_KEY  = 'vtEUNe78i4FlIaHTNj4oKy8wZWq5a78Z'

client = AipOcr(APP_ID, API_KEY, SECRET_KEY)

def get_file_content(file):
    with open(file, 'rb') as fp:
        return fp.read()

def img_to_str(image_path):
    image = get_file_content(image_path)
    result = client.basicGeneral(image)
    if 'words_result' in result:
        return u'\n'.join([w['words'] for w in result['words_result']])

print(img_to_str(image_path='/Users/general/Pictures/ocr-img.png'))</string>
						<key>TimeOutAbortsMacro</key>
						<true/>
						<key>TrimResults</key>
						<true/>
						<key>TrimResultsNew</key>
						<true/>
						<key>UseText</key>
						<true/>
					</dict>
					<dict>
						<key>ActionName</key>
						<string>发送通知</string>
						<key>MacroActionType</key>
						<string>Notification</string>
						<key>SoundName</key>
						<string>DefaultSoundName</string>
						<key>Subtitle</key>
						<string>OCR 结果已经复制到剪贴板</string>
						<key>Text</key>
						<string></string>
						<key>Title</key>
						<string>%ExecutingMacro%</string>
					</dict>
					<dict>
						<key>ActionName</key>
						<string>删除临时保存文件</string>
						<key>Destination</key>
						<string></string>
						<key>MacroActionType</key>
						<string>File</string>
						<key>Operation</key>
						<string>Delete</string>
						<key>Source</key>
						<string>~/Pictures/ocr-img.png</string>
					</dict>
				</array>
				<key>CreationDate</key>
				<real>574520593.86887801</real>
				<key>ModificationDate</key>
				<real>652351966.68706501</real>
				<key>Name</key>
				<string>OCR 截图</string>
				<key>Triggers</key>
				<array>
					<dict>
						<key>FireType</key>
						<string>Pressed</string>
						<key>KeyCode</key>
						<integer>0</integer>
						<key>MacroTriggerType</key>
						<string>HotKey</string>
						<key>Modifiers</key>
						<integer>4864</integer>
					</dict>
				</array>
				<key>UID</key>
				<string>980745EF-E465-4E5E-A2BB-D7897061305F</string>
			</dict>
		</array>
		<key>Name</key>
		<string>Sharing</string>
		<key>ToggleMacroUID</key>
		<string>FE8BFC0D-71E7-4F36-BB3D-2CEFB7C3F5B8</string>
		<key>UID</key>
		<string>73EB426A-F550-4527-8093-6240E0DCCCCA</string>
	</dict>
</array>
</plist>

```


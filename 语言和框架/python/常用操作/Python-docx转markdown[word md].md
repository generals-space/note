# Python-docx转markdown[word md]

参考文章

1. [Python将DOCX转换为markdown文件](https://blog.csdn.net/weixin_43431593/article/details/105185702)
    - 先将 docx 转成 html, 再转成 markdown(不过代码有问题)
    - 转成的 html 中, image 图片以 base64 的格式内置在 html 代码中
2. [纯Python 实现 Word 文档转换 Markdown](https://my.oschina.net/u/3454592/blog/4829127)
    - 有提到 pandoc
    - 同样是先转成 html, 再转成 markdown
    - 转成的 html 中, image 图片同样以 base64 的格式内置在 html 代码中, 但是提供了将图片写入本地文件的方法

pythono: 3.6

```
markdownify=0.6.5
mammoth=1.4.15
```

```py
#!/bin/python3

import os
import sys
import time
import hashlib
import pathlib

import mammoth
from markdownify import markdownify

md_path = 'test.md'

img_path = './src'

def init_img_path():
    '''
    创建图片存储目录
    '''
    if not pathlib.Path(img_path).exists(): 
        ## parents=True 表示如果目标目录的父级目录不存在的情况下自动创建, 
        ## 同 shell 中的 mkdir -p
        pathlib.Path(img_path).mkdir(parents = True)

def convert_img(img):
    '''
    对 word 文档中的图片进行处理, 将图片名直接转换成 md5 字符串
    '''
    img_suffix = img.content_type.split('/')[1]

    ## 这里只能用 with 语句, 如果改成 img_bytes = img.open() 会报错.
    with img.open() as img_bytes:
        img_byte_cnt = img_bytes.read()
        md5 = hashlib.md5()
        md5.update(img_byte_cnt)
        md5_str = md5.hexdigest()
        path_file = '{}/{}.{}'.format(img_path, md5_str, img_suffix)
        f = open(path_file, 'wb')
        f.write(img_byte_cnt)
        f.close()

    return {'src':path_file}

def doc2md(docx_path:str):
    init_img_path()
    ## 读取 Word 文件(二进制方式)
    docx_file = open(docx_path ,'rb')
    ## 获取无后缀的文件名称, 这里是为了防止文件名中包含多个点号, 所以没有直接使用[1]进行选择.
    docx_name = '.'.join(docx_path.split('.')[:-1])
    ## 转化 Word 文档为 HTML
    ## mammoth 有 convert_to_markdown() 方法
    result = mammoth.convert_to_html(docx_file, convert_image = mammoth.images.img_element(convert_img))
    docx_file.close()

    ## 获取 HTML 内容
    html_cnt = result.value
    ## messages = result.messages
    ## html_path = 'test.html'
    ## html_file = open(html_path, 'w', encoding='utf-8')
    ## html_file.write(html_cnt)
    ## html_file.close()

    ## 转化 HTML 为 Markdown
    md_cnt = markdownify(html_cnt, heading_style = 'ATX')
    md_file = open(docx_name + '.md', 'w', encoding='utf-8')
    md_file.write(md_cnt)

    md_file.close()

if __name__ == '__main__':
    ## python k2file.py /etc/kubernetes/admin.conf 的 argv 中不包含 `python`
    if len(sys.argv) == 1: 
        print("请指定目标文件路径")
        sys.exit(-1)
    doc2md(sys.argv[1])

```

使用方法

```
python3 main.py xxx.docx
```

会在当前目录生成同名的`.md`文件.

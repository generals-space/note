# 使用pandoc进行markdown与word的格式转换[md docx]

```
d run -it --name pandoc --rm -v /home/playground/pandoc:/data pandoc/core sh
```

## markdown -> word

没找到合适的, markdown转成word格式都没了. 有人建议直接用编辑器的预览功能, 然后拷贝到word中.

不过我目前没有这个需求, 先不考虑.

## word -> markdown

```
pandoc -f docx -t markdown -o test.md test.docx
```

如果word中包含图片, 可以添加`--extract-media`选项将其保存到本地, 生成的markdown文件中会以相对路径的格式引用.

```
pandoc -f docx -t markdown --extract-media='.' -o test.md test.docx
```


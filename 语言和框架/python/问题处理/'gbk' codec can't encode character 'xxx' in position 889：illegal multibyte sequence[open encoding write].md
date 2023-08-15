# 'gbk' codec can't encode character 'xxx' in position 889: illegal multibyte sequence

参考文章

1. [强推！！！解决UnicodeEncodeError: ‘gbk‘ codec can‘t encode character.....: illegal multibyte](https://blog.csdn.net/m0_37772653/article/details/119783057)

## 问题描述

```
code_file = open(file_path_obj, 'w+')
code_file.write('xxx')
```

如果写入的文件内容中包含一些特殊字符, 则在`write()`时会报如下错误

```
'gbk' codec can't encode character '\xb5' in position 889: illegal multibyte sequence
```

## 解决方法

```
code_file = open(file_path_obj, 'w+', encoding='utf-8')
```

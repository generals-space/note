# json.dumps使用[中文 unicode 缩进]

参考文章

1. [python json.dumps 中文编码](https://www.cnblogs.com/shiju/p/9511916.html)

`json.dump(object)`是输出到文件, 而`json.dumps(object)`是输出为字符串.

## 中文

```py
>>> import json
>>> dic = {'country': '中国'}
>>> json.dumps(dic)
'{"country": "\\u4e2d\\u56fd"}'
>>> json.dumps(dic, ensure_ascii=False)
'{"country": "中国"}'
```

## 缩进 pretty

```py
json.dumps(dic, ensure_ascii=False, indent=4)
```

这个在交互式命令行看不出来, 在程序里可以使用.

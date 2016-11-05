# Grep应用方法总结

## 1. 显示匹配行的前后几行

- -A num: after, 显示目标匹配行的后num行

- -B num: before, 显示目标匹配行的前num行

- -C num: 显示目标匹配行的前num行与后num行

```
## 显示file文件里匹配string字串那行及其上下5行
grep -C 5 string file
## 显示string及前5行
grep -B 5 string file
## 显示string及后5行
grep -A 5 string file
```

## 2. 只显示匹配文件名

原始的`grep`命令查找的结果如下.

```
$ grep -ri 'cc' ./*
./aa.txt:ccc
./bb/bb.txt:cc
```

使用`-l`参数可以单独输出文件路径而不再输出匹配行的内容.

```
$ grep -ril 'cc' ./*
./aa.txt
./bb/bb.txt
```

注意: 使用了`-l`, 则`-n`将失效, 不再输出匹配行号.
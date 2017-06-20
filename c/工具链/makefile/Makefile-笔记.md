# 笔记-Makefile

## 1.

同时生成多个可执行文件时，需要用`all`标签预先声明，否则make只会默认生成第一个可执行文件

示例代码

```makefile
CC = g++
CXXFLAGS = -std=c++11
lexicalObj = scanner.o lexical.o
parserObj = scanner.o parser.o compiler.o
#you need to declare a 'all' to show all excuted file
#or you will just get the first excuted target
all : compiler lexical
#lexical
lexical : $(lexicalObj)
	$(CC) $(CXXFLAGS) -g $^ -o $@
lexical.o : lexical.cpp 
	$(CC) $(CXXFLAGS) -c $< -o $@
#parser
compiler : $(parserObj)
	$(CC) $(CXXFLAGS) -g $^ -o $@
scanner.o : scanner.cpp scanner.h
	$(CC) $(CXXFLAGS) -c $< -o $@
parser.o : parser.cpp parser.h
	$(CC) $(CXXFLAGS) -c $< -o $@
compiler.o : compiler.cpp
	$(CC) $(CXXFLAGS) -c $< -o $@
#while cleaning, do not rm a file twice
#or it will get an error
clean :
	rm lexical lexical.o compiler $(parserObj)
```

## 2.

`clean`标签中删除中间文件时，不可删除一个文件两次，否则会出错。注意变量所代表的中间文件的重复。
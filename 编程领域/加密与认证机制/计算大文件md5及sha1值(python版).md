# 计算大文件md5及sha1值(python版)

参考文章

1. [python计算大文件的md5、sha1值](https://segmentfault.com/a/1190000002978881)

2. [超大文件如何计算md5？](https://www.zhihu.com/question/34364469)

3. [python中使用md5和sha1加密](https://blog.csdn.net/a599174211/article/details/83578294)

对于大文件, 计算其md5值时如果把整个文件读取到内存再计算是很不明智的, 分段计算才是好办法, 因为md5和sha1等对象有`update()`方法, 可以实现这样的操作.

```py
#!/usr/bin/python3

import hashlib
import base64

## hashlib中只有md5和shaXXX系的加密函数...没有其他可用的.

## 计算文件的sha1值
def SHA1FileWithName(fineName, block_size=64 * 1024):
    with open(fineName, 'rb') as f:
        sha1 = hashlib.sha1()
        while True:
            data = f.read(block_size)
            if not data: break
            sha1.update(data)
        ## result = base64.b64encode(sha1.digest())
        result = sha1.hexdigest() ## ## 这将得到传统的sha1字段串
        return result

## 计算文件的md5值
def MD5FileWithName(fineName, block_size=64 * 1024):
    with open(fineName, 'rb') as f:
        md5 = hashlib.md5()
        while True:
            data = f.read(block_size)
            if not data: break
            md5.update(data) ## 新增文件分段, 更新摘要
        ## result = base64.b64encode(md5.digest())
        result = md5.hexdigest() ## 这将得到传统的md5字段串
        return result

def MD5StringValue(string):
    md5 = hashlib.md5()
    md5.update(string.encode('utf-8'))
    result = md5.hexdigest()
    return result

def SHA1StringValue(string):
    sha1 = hashlib.sha1()
    sha1.update(string.encode('utf-8'))
    result = sha1.hexdigest()
    return result

if __name__ == '__main__':
    filepath = '/Users/general/Movies/Lost Ark - Open Beta CG Trailer - PC - F2P - KR.mp4'
    sha1code = SHA1FileWithName(filepath)
    print(sha1code)
    md5code = MD5FileWithName(filepath)
    print(md5code)
    print(MD5StringValue('hello world'))
    print(SHA1StringValue('hello world'))
```
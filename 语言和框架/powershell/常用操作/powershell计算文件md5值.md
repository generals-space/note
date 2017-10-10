# powershell计算文件md5值

<!tags!>: <!md5!>

`get-filehash -algorithm 算法类型 目标文件`

算法类型包括:

SHA1 | SHA256 | SHA384 | SHA512 | MACTripleDES | MD5 | RIPEMD160

例如

```ps
$ set-content test 'hello world'

$ cat test
hello world

$ get-filehash -algorithm md5 test

Algorithm       Hash                                                                   Path
---------       ----                                                                   ----
MD5             A0F2A3C1DCD5B1CAC71BF0C03F2FF1BD                                       C:\Users\general\Downloads\test
```
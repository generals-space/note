
`os.path.getsize(path)`: 返回文件大小(单位为**字节**)，如果文件不存在就返回错误.

`os.path.exists(path)`: 判断`path`表示的路径是否存在.

`os.path.isfile(path)`: 判断是否为文件.

`os.path.isdir(path)`: 判断是否为文件.

------

## 文件操作

`os.remove(file_path)`: 删除`file_path`所表示的文件

`os.removedirs(dir_path)`: 删除`file_path`所表示的目录, 但删除非空目录不好使

`os.makedirs(path)`: 创建空目录

`shutil.copyfile(src, dst)`: 拷贝文件, `src`和`dst`都是(可包含路径的)文件名. 不可以是目录, 否则会报错. 

`shutil.rmtree(path)`: 可删除非空目录, 即递归删除

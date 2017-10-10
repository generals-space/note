# Django应用总结

## 1. 模板转义


## 3. Django获取前端checkbox值

多选情况下, 前端代码为

```html
<form action="" method="POST">
    <input type="checkbox" value="1" name="check_box_list"/>苹果
    <input type="checkbox" value="2" name="check_box_list"/>梨
    <input type="checkbox" value="3" name="check_box_list"/>杏
    <input type="checkbox" value="4" name="check_box_list"/>桃子
    <input type="submit" value="提交">
</form>
```

后端django代码

```python
if request.method == 'POST':
    ## check_box_list是一个list对象, 选中的值会出现在里面, 如[2, 4]
    check_box_list = request.REQUEST.getlist('notuseldap')
```
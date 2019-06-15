# gorm-Scan与Pluck查询

## Scan查询

```go
type User struct {
	Name string
	Age uint8
}
```

一般与`Select()`配合使用, 当我们只想要查询某张表的指定字段时, 其他字段是没有用处的, 那么常规的`Find(&users)`其实有点得不偿失的感觉, 尤其是当这张表中的字段很多时, 结果列表可能占用较大内存. 

此时我们可以使用`Scan()`方法, 仅将我们需要的列放入结果中即可.

```go
type NameResult struct{
	Name string
}
nameResults := []*NameResult{}
// nameResults := []*User{} // 其实Scan()可以将结果填入User数组对象
db.Model(&User{}).Select("name").Scan(&nameResults)
```

> 注意: 如果使用`First()`或`Find()`去查询`&nameResults`的值什么也不会得到, 猜测是因为与`Model()`的参数类型不同. 就算没有Model, 也会因为数据库中根本不存在`name_results`表而查询失败.

> 另外, `Scan()`也可以配合`Raw()`方法执行原生sql使用.

## Pluck查询

与Scan类似...上面的代码中我只想查单个列, 还要新建一个结构体, 感觉不值得, 因为单个列的最终结果其实应该是一个字符串列表, `Pluck()`就是将目标结果赋值给一个列表的函数.

```go
	results := []string{}
	db.Model(&model.Manufacturer{}).Pluck("name", &results)
```

这样更是简单, `Scan`更适合查询指定列但是数量大于1列的情况.

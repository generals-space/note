# gin模板应用

参考文章

1. [golang 模板(template)的常用基本语法](https://blog.csdn.net/zhang_guyuan/article/details/76549436)

2. [Go templates with eq and index](https://stackoverflow.com/questions/33282061/go-templates-with-eq-and-index)

3. [Go template: can't evaluate field X in type Y (X not part of Y but stuck in a {{range}} loop)](https://stackoverflow.com/questions/43263280/go-template-cant-evaluate-field-x-in-type-y-x-not-part-of-y-but-stuck-in-a)

4. [Go语言中使模板引擎的语法](https://zhuanlan.zhihu.com/p/50397556)

	- 模板中的可用函数

## 1. if判断中的复杂运算

变量`Categories`内容如下

```go
	Categories := []*frontend.CategoryAndBooks{
        &frontend.CategoryAndBooks{
			Category: &frontend.CategoryInfo{
				Title: categoryJSON.Title,
				URL:   categoryURL,
			},
		},
    }

```

```html
<meta name="description" content="{{range $index, $item := .Categories}}{{$item.Category.Title}}小说{{if lt (add $index 1) (len $.Categories)}}，{{else}}。{{end}}{{end}}" />
```

需要注意的是`if`判断中的两个值运算, 需要用`()`包裹, 可见参考文章2.

另外还有`$.Categories`, 在`range`循环中, 无法再直接通过`.Categories`得到外层变量, 根节点需要通过`$`来指代, 这一点可见参考文章3.
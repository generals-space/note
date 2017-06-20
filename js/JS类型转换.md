# JS类型转换

<!tags!>: <!js!> <!类型转换!>

[JS 对象（Object）和字符串（String）互转](http://blog.csdn.net/starrexstar/article/details/8083259/)

## 1. JS字典列表和字符串相互转换

利用原生JSON对象，将对象转为字符串

```js
var jsObj = {};
jsObj.testArray = [1,2,3,4,5];
jsObj.name = 'CSS3';
jsObj.date = '8 May, 2011';
var str = JSON.stringify(jsObj);
console.log(str);
```

从JSON字符串格式化为对象

```js
var jsObj = {};
jsObj.testArray = [1,2,3,4,5];
jsObj.name = 'CSS3';
jsObj.date = '8 May, 2011';
var str = JSON.stringify(jsObj);
console.log(str);
var str1 = JSON.parse(str);
console.log(str1);
```
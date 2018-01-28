
```js
(function abc(a){
    console.log(a);
})(2); // 这里传入的2就是匿名函数的形参
```

```js
var abc = function(a){
	console.log(a);
}(3);
```

```js
!function abc(a){
	console.log(a);
}(3);
```
# JS-阻塞式sleep

```js
function sleep(d){  
    var t = Date.now();
    while(Date.now - t <= d);  
}
// sleep一秒钟
sleep(1000);
```
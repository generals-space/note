参考文章

1. [大批量DOM操作如何提升性能](https://segmentfault.com/q/1010000011813044)

2. [【优化】分时加载](https://www.cnblogs.com/bluedream2009/archive/2010/03/16/1687095.html)

几点说明：

1. timedChunk 函数，里面的 50ms 来自 Response Time Overview 中的调查结果：100ms 内的响应能让用户感觉非常流畅。50ms 是 Nicholas 针对 JavaScript 得出的最佳经验值。

2. setTimeout 延时 25ms, 是因为浏览器的时间分辨率问题。25ms 可以保证主流浏览器都顺畅（有喘息的机会去更新 UI）。

3. 上面的实例，传统方式加载会让浏览器在加载数据期间，无法更新界面和响应任何操作。采用分时加载，则可以让浏览器始终保持可响应状态，提升界面流畅性和用户体验。

```html
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<title> 分时加载 </title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<style>
button { width:200px; height:50px; }
#lay {margin:0 auto; width: 670px; }
#userlist { width:502px; border:1px solid #8C8C8C; margin-left:50px; margin-top:10px; height:30px; border-collapse:collapse; font-size:12px; }
#userlist thead td { height:25px; border:1px solid #8C8C8C; text-align:center; background-color:#00CC00; }
#userlist tbody td { height:25px; border:1px solid #8C8C8C; text-align:center; }
</style>
</head>
<body>
    <div id="lay">
        <button id="resetBtn">重置</button>
        <button id="tradBtn">传统加载</button>
        <button id="timeBtn">分时加载</button>
        <div style="width:500px; border:1px solid #000; background-color:#FFF; margin-left:50px; margin-top:10px; ">
            <div id="processBar" style="width:0; height:35px; background-color:#003366; line-height:35px; color:#FFF; ">
            </div>
        </div>
        <table id="userlist">
            <thead><tr><td>用户列表</td></tr></thead>
            <tbody></tbody>
        </table>
    </div>

<script>
var loadDemo = function() {
    var $ = function(id) {
        return document.getElementById(id);
    },
    resetBtn = $('resetBtn'), 
    tradBtn = $('tradBtn'), 
    timeBtn = $('timeBtn'), 
    progress = $('processBar'), 
    userlist = $('userlist'), 
    tbody = userlist.tBodies[0], 
    i, n = 1, 
    JSON_DATA = [];
 
    // 创建用户列表
    function addItem(data) {
        var tr = document.createElement('tr');
        var td = document.createElement('td');
        td.appendChild(document.createTextNode(data[0]));
        tr.appendChild(td);
        tbody.appendChild(tr);
        if(++n % 10 == 0) progressing();
    }
    // 加载进度
    function progressing() {
        progress.style.width = (progress.offsetWidth + 1) + 'px';
    }
    // 清空用户列表
    function reset() {
        resetBtn.disabled = true;
        tradBtn.disabled = false;
        timeBtn.disabled = false;
        if(window.ActiveXObject) { // IE
            var temp = document.createElement('div');
            temp.innerHTML = '<table><tbody></tbody></table>';
            userlist.replaceChild(temp.firstChild.tBodies[0], tbody);
            tbody = userlist.tBodies[0];
        } else {
            tbody.innerHTML = '';
        }
        progress.innerHTML = '';
        progress.style.width = 0;
        n = 1;
    }
    // 分时函数[*****]
    function timedChunk(items, process, context, callback) {
        // 拷贝副本, 之后的操作中每处理一个成员就从中移除
        var todo = items.concat(), delay = 25;
        setTimeout(function() {
            var start = +new Date();
            do {
                process.call(context, todo.shift());
            } while (todo.length > 0 && (+new Date() - start < 50))

            if(todo.length > 0) {
                setTimeout(arguments.callee, 25);
            } else if(callback) {
                callback();
            }
        }, delay);
    }
    // 分时加载
    function testTimed() {
        resetBtn.disabled = true;
        tradBtn.disabled = true;
        timeBtn.disabled = true;
        var start = +new Date(), end;
        timedChunk(JSON_DATA, addItem, null, function() {
            resetBtn.disabled = false;
            end = +new Date();
            progress.innerHTML = (end - start) + ' ms';
        });
    }
    // 传统加载
    function testTrad() {
        resetBtn.disabled = true;
        tradBtn.disabled = true;
        timeBtn.disabled = true;
        var start = +new Date(), end;
        for(var i = 0, len = JSON_DATA.length; i < len; i++) {
            addItem(JSON_DATA[i]);
        }
        end = +new Date();
　　　　 resetBtn.disabled = false;
        progress.innerHTML = (end - start) + ' ms';
    }
 
    return {
        run: function() {
            // 模拟用户数据
            for(i = 1; i <= 5000; i++) {
                JSON_DATA.push(['用户名称' + i]);
            }
            resetBtn.onclick = function() { reset(); }
            tradBtn.onclick = function() { testTrad(); }
            timeBtn.onclick = function() { testTimed(); }
        }
    }
}();
 
loadDemo.run();
 
</script>
</body>
</html>
```

------

其中`timedChunk`函数中用到了`arguments.callee`对象, 但是在严格模式下这个对象不允许被访问. 可以写成如下形式

```js
// 分时函数[*****]
function timedChunk(items, process, context, callback) {
    // 拷贝副本, 之后的操作中每处理一个成员就从中移除
    var todo = items.concat(), delay = 25;
    function _timedChunk(){
        var start = +new Date();
        do {
            process.call(context, todo.shift());
        } while (todo.length > 0 && (+new Date() - start < 50))

        if(todo.length > 0) {
            setTimeout(_timedChunk, 25);
        } else if(callback) {
            callback();
        }
    }
    setTimeout(_timedChunk, delay);
}
```
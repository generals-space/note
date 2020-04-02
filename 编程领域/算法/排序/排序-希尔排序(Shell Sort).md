希尔排序是对简单插入入排序的优化, 基于以下两个认识：

1. 数据量较小时插入排序速度较快，因为n和n方差距很小；
2. 数据基本有序时插入排序效率很高，因为比较和移动的数据量少。

希望排序中的核心变量就是`gap`, ta就像梳子的齿一样, 开始时使用空隙较为大的齿梳理一遍, 然后慢慢细致, 最后必须要保证`gap`值为1. `gap`值为1时就相当于是简单插入排序了.

查看参考文章中4中对希尔排序的配图就会很容易懂其工作原理.

至于优化效果, 简单插入排序的时间复杂度为n方. 假设n==10, 那复杂度就是100.

希尔排序挑选gap值分别为: 5, 2, 1这3个值来运行, 消耗时间为1 + 4 + 25, 远小于100...希望我没理解错.

```py
import math

arr = [5, 38, 15, 48, 44, 3, 36, 26, 50, 27, 2, 46, 4, 19, 47]

def insert_sort(arr, length, gap):
    for i in range(gap, length):
        ## 以第一个元素为已排序过的元素, 遍历之后的元素向前/向后插入.
        current = arr[i]
        preIndex = i - gap
        ## 遍历已排过序的部分(从后往前), 
        ## 一边遍历, 一边将比current大的成员向后移动, 
        ## 但不必遍历全部, 找到比当前索引i处数值current更小的成员即可插入.
        while preIndex >= 0 and arr[preIndex] > current:
            arr[preIndex + gap] = arr[preIndex]
            preIndex -= gap
        ## 这里比较容易出错, 注意加1
        arr[preIndex + gap] = current
    return arr

def shell_sort(arr):
    length = len(arr)
    gap = math.floor(length/2)
    while True:
        if gap == 0: break
        insert_sort(arr, length, gap)
        gap = math.floor(gap/2)
    return arr

print(arr)
arr = shell_sort(arr)
print(arr)

```


```js
var arr = [5, 38, 15, 48, 44, 3, 36, 26, 50, 27, 2, 46, 4, 19, 47];

console.log(arr);
arr = shellSort(arr);
console.log(arr);

// 常规的插入排序, gap值为1
function insertSort(arr, len, gap) {
    for (let i = gap; i < len; i++) {
        let current = arr[i];
        let preIndex = i - gap;
        while (preIndex >= 0 && current < arr[preIndex]) {
            arr[preIndex + gap] = arr[preIndex];
            preIndex = preIndex - gap;
        }
        arr[preIndex + gap] = current;
    }
}

function shellSort(arr) {
    let len = arr.length;
    // gap趋向于越来越小, 直到1
    for (let gap = Math.floor(len / 2); gap > 0; gap = Math.floor(gap / 2)) {
        insertSort(arr, len, gap);
    }
    return arr;
}

```
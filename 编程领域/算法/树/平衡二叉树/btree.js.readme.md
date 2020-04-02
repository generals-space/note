# js版平衡二叉村总结

## 1. js版本新节点的声明与赋值, 和空节点的判断, 借鉴了C语言版本.

`var node = {}`对应了C语言中声明为`BTNode *root = NULL`, 作为空节点;

`if (!Object.keys(node).length)`则对应 C语言中的`if (node == NULL)`, 作为判断条件;

即js中的{}作用等同于C中的NULL.

之前的做法

```js
function newNode() {
    return {
        left: undefined,
        right: undefined,
        value: NaN,
        height: 1,
    }
}
```

然后用在`insert()`中

```js
if (!node.value) {
    node.value = val;
    node.left = newNode();
    node.right = newNode();
    return node;
}
```

但这样的话, 叶子节点上的left和right也会有空对象存在(只不过value为NaN罢了).

```
                 2
              /     \
            /         \ 
          /             \
        1                 4
      /   \             /   \
    /       \         /       \
  NaN       NaN      3         5
                   /   \     /   \
                 NaN   NaN NaN   NaN
```

那么在遍历的时候也需要判断当前节点的value值是否为NaN才行, 否则会输出这样的值.

...绝对不能这么做.

## 2. height()函数中的节点判空操作

## 3. LR/RL中第1步的操作分别针对node.left和node.right, 而不是node本身.

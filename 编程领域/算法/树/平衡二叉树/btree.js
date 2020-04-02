function newNode(node) {
    node.left = {};
    node.right = {};
    node.value = NaN;
    node.height = 1;
    return node;
}

/**
 * 判断目标节点node是否为空, 作用等同于C语言版本中的 if(Node == NULL)
 * `if ({}) console.log(123);` 将会输出123, 
 * 所以无法使用 node ? node.height : 0 三元表达式.
 * @param {}} node 
 */
function isEmpty(node) {
    return !(Object.keys(node).length > 0);
}

/**
 * 获取目标子树的高度, 其实是直接返回node的height成员.
 * @param {*} node 
 */
function height(node) {
    if (isEmpty(node)) return 0;
    return node.height;
}

function max(a, b) {
    return a > b ? a : b;
}
/**
 * 获取node节点的平衡因子, 即左右两个子树的高度差.
 * 大于0则表示左子树比较深, 反之则右子树比较深.
 * @param {*} node 
 * @returns        返回整型数值, 可能为负值.
 */
function getBlance(node) {
    // node没有左/右子节点并不算特殊情况, 这么做只是避免程序执行出错.
    let left = node.left ? height(node.left) : 0;
    let right = node.right ? height(node.right) : 0;
    return left - right;
}

/**
 * 
 * @param {*} node 
 */
function ll_rotate(node) {
    let tmp = node.left;
    node.left = tmp.right;
    tmp.right = node;

    // 旋转完成, 重设两个变换节点的height成员值.
    node.height = max(height(node.left), height(node.right)) + 1;
    tmp.height = max(height(tmp.left), height(tmp.right)) + 1;
    return tmp;
}

/**
 * 
 * @param {*} node 
 */
function rr_rotate(node) {
    let tmp = node.right;
    node.right = tmp.left;
    tmp.left = node;

    // 旋转完成, 重设两个变换节点的height成员值.
    node.height = max(height(node.left), height(node.right)) + 1;
    tmp.height = max(height(tmp.left), height(tmp.right)) + 1;

    return tmp;
}

/**
 * 将新节点val插入node所表示的根节点的二叉平衡树中.
 * @param {*} root 目标树的根节点
 * @param {*} val  新插入的节点的值
 * @returns        插入过程中可能会经过调整, 此函数会返回该子树新的根节点.
 */
function insert(node, val) {
    // 如果node.value为空, 说明该树为空, 把val作为根节点返回.
    // C语言中以node == NULL为判断条件, 这里用value成员的NaN特殊值做条件.
    // 如果是python呢? python没有结构体, 可以用class模拟吧???
    if (isEmpty(node)) {
        console.log(node);
        node = newNode(node);
        node.value = val;
        return node;
    }

    if (val < node.value) {
        node.left = insert(node.left, val)
    } else if (val > node.value) {
        node.right = insert(node.right, val)
    } else {
        // 二叉平衡树中不会出现相同的节点值, 所以这种情况直接返回.
        return node;
    }

    // console.log(node);
    // 插入完成, 有可能树的层高会发生变化, 所以需要修正node.height值.
    node.height = max(height(node.left), height(node.right)) + 1;
    // console.log('====', node);
    // 插入完成后, 可能二叉树已经失去平衡, 所以需要调整.
    balance = getBlance(node);
    if (balance > 1 && node.left.value > val) return ll_rotate(node);
    if (balance < -1 && node.right.value < val) return rr_rotate(node);
    if (balance > 1 && node.left.value < val) {
        // 特别注意: 这里RR的目标是node.left, node本身的位置是不会变的.
        node.left = rr_rotate(node.left);
        node = ll_rotate(node);
        return node;
    }
    if (balance < -1 && node.right.value > val) {
        // 特别注意: 这里LL的目标是node.right, node本身的位置是不会变的.
        node.right = ll_rotate(node.right);
        return rr_rotate(node);
    }
    return node;
}

function preOrder(node) {
    // 防止node为{}时, console.log()输出 undefined.
    if(isEmpty(node)) return;
    console.log(node.value);
    if (node.left) preOrder(node.left);
    if (node.right) preOrder(node.right);
    return;
}

var list = [9, 5, 10, 0, 6, 11, -1, 1, 2];
var root = {};
list.forEach((val, key) => {
    console.log(val);
    root = insert(root, val);
});

preOrder(root);

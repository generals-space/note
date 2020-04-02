#!/usr/bin/python
'''
已知前序与后序遍历序列, 还原二叉树对象.
'''

class TreeNode:
    '''
    节点对象
    '''
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

def rebuild_btree(pre, tin):
    '''
    @param pre: 前序遍历序列

    @param tin: 中序遍历序列

    需要注意的是, 在递归过程中, 传入的pre和tin参数长度总是相同的.
    因为ta们是同一个树(或是子树)的不同遍历序列.
    '''
    if len(pre) == 0: return None

    ## 前序遍历的第一个成员为根节点.
    root = pre[0]
    rootNode = TreeNode(root)
    ## 结束递归的标识
    if len(pre) == 1: return rootNode

    ## 根节点下左右子树的中序遍历.
    tin_l, tin_r = [], []
    for i, v in enumerate(tin):
        if v == root: 
            tin_l, tin_r = tin[0:i], tin[i+1:]
            break

    ## 根节点下左右子树的前序遍历.
    pre_l, pre_r = [], []
    for i, v in enumerate(pre):
        ## 跳过第一个根节点(此时pre长度一定大于1)
        if i == 0: continue
        ## 下面的代码有bug, 因为需要考虑pre[1:]中的成员全在tin_l中的情况,
        ## 即当前root节点下, 只有左子树没有右子树的情况.
        ## if v not in tin_l:
        ##     pre_l, pre_r = pre[1:i], pre[i:]
        if v in tin_l:
            pre_l.append(v)
        else:
            pre_r = pre[i:]
            break

    ## 如何结束递归?
    ## 本函数最终一定会return rootNode的.
    rootNode.left = rebuild_btree(pre_l, tin_l)
    rootNode.right = rebuild_btree(pre_r, tin_r)
    return rootNode

pre = [1,2,4,7,3,5,6,8]
tin = [4,7,2,1,5,3,8,6]

rebuild_btree(pre, tin)

## 此函数在牛客网上经过了测试, 可以不用怀疑其正确性.

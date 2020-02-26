
#include<stdio.h>
#include<stdlib.h>

typedef struct Node
{
    int key;
    struct Node *left;
    struct Node *right;
    // 表示当前节点所处的深度(初始根节点为1), 主要用于获取左右子树的高度差,
    // 判断是否处于不平衡状态, 左右子树哪一边更深.
    int height;
}BTNode;

int max(int a, int b);

// @function: 返回目标node的height成员值, 
// 程序在构造与维护(包括插入和删除操作)平衡二叉树时维护此值.
int height(struct Node *N)
{
    if (N == NULL) return 0;
    return N->height;
}

// @function: 取最大值工具函数.
int max(int a, int b)
{
    return (a > b) ? a : b;
}

BTNode* newNode(int key)
{
    struct Node* node = (BTNode*)malloc(sizeof(struct Node));
    node->key = key;
    node->left = NULL;
    node->right = NULL;
    node->height = 1;
    return(node);
}

/*
    调用LL调整的两种情况:

    1. LL
                            LL
                4           =>          4
              /   \                   /   \
    node ->  3     5                 2     5
           /                       /   \
          2                key -> 1     3 <- node
        /                         
       1 <- key                   

    2. RL
                             LL(node.right)                    RR            
          2                  =>             2                  =>               2
        /   \                             /   \                               /   \
       1     4  <- node                  1     4  <- node                    1     3  <- node
               \                                 \                               /   \
                5                         key ->  3                   node ->  4     5
              /                                     \                           
             3  <- key                               5                          
*/
BTNode* ll_rotate(BTNode* node)
{
    BTNode *tmp = node->left;
    node->left = tmp->right;
    tmp->right = node;
    // 旋转完成, 补全height成员值, 补全的操作还是很常规的.
    node->height = max(height(node->left), height(node->right)) + 1;
    tmp->height = max(height(tmp->left), height(tmp->right)) + 1;

    return tmp;
}

/*
    调用RR调整的两种情况:
    
    1. RR
                             RR                         
           4                 =>          4              
         /   \                         /   \            
        3     5  <-node               3     5  <- node  
                \                         /   \         
                 6                       6     7  <- key
                   \                                    
             key -> 7                                   
    
    2. LR
                            RR(node.left)              LL
                   4        =>             4           =>               4
                 /   \                   /   \                        /   \
      node ->  3      5        node ->  3     5              key ->  2     5
              /                       /                            /   \
             1                       2  <- key                    1     3  <- node
               \                   /                   
                2  <- key         1                    
*/
BTNode* rr_rotate(BTNode* node)
{
    BTNode *tmp = node->right;
    node->right = tmp->left;
    tmp->left = node;
    // 旋转完成, 补全height成员值.
    node->height = max(height(node->left), height(node->right)) + 1;
    tmp->height = max(height(tmp->left), height(tmp->right)) + 1;

    return tmp;
}

// @function: 返回目标节点左右子树的高度差(所谓的平衡因子), 
// 大于0则表示左子树比较深, 反之则右子树比较深.
// 主调函数用此函数判断该节点下的二叉树是否处于平衡.
int getBalance(BTNode* N)
{
    // 其实N==NULL不算特殊情况, 只不过不这么写程序执行会出错.
    if (N == NULL) return 0;
    return height(N->left) - height(N->right);
}

// @function:
// 使用insert()插入节点构造的树不会出现太离谱的状况, 因为一旦出现不平衡就会被修正.
// 当低层子树经过旋转调整后, 可能会引起高层子树重新陷入不平衡, 所以还要再次检测并修正.
// @param node: 目标树的根节点指针.
// @param key:  该节点的值.
// @reuturn:    总是返回node子树在插入新节点key后的新的根节点.
BTNode* insert(BTNode* node, int key)
{
    // 空树时直接当根节点然后返回.
    // 但还要注意一点, 整个程序中只有这一处调用了newNode()函数.
    // 这也是下面比较key与node左右节点大小的目的, 先找到合适的位置, 
    // 再创建新节点然后插入.
    if (node == NULL) return newNode(key);

    if (key < node->key)
        node->left = insert(node->left, key);
    else if (key > node->key)
        node->right = insert(node->right, key);
    else
        // ...遇到与已有节点数值节点的节点, 直接返回, 不添加.
        // 貌似是因为平衡二叉树的概念需要各节点中的值不相同.
        return node;

    // 插入完成后, 有可能树的层高会发生变化, 更新node->height的值.
    node->height = 1 + max(height(node->left), height(node->right));

    int balance = getBalance(node);

    // 注意: 下面4种情况的if条件中, 除了balance, 另外的是要与node的左/右节点相比较.
    // LL型
    if (balance > 1 && key < node->left->key) return ll_rotate(node);
    // RR型
    if (balance < -1 && key > node->right->key) return rr_rotate(node);
    // LR型
    if (balance > 1 && key > node->left->key)
    {
        // 这里RR的目标是node->left, node本身的位置是不会变的.
        node->left = rr_rotate(node->left);
        return ll_rotate(node);
    }
    // RL型
    if (balance < -1 && key < node->right->key)
    {
        // 这里LL的目标是node->right, node本身的位置是不会变的.
        node->right = ll_rotate(node->right);
        return rr_rotate(node);
    }

    return node;
}

BTNode * minValueNode(BTNode* node)
{
    BTNode* current = node;
    while (current->left != NULL) current = current->left;
    return current;
}

BTNode* deleteNode(BTNode* root, int key)
{
    if (root == NULL) return root;

    if (key < root->key) 
        root->left = deleteNode(root->left, key);
    else if (key > root->key)
        root->right = deleteNode(root->right, key);
    else
    {
        if ((root->left == NULL) || (root->right == NULL))
        {
            BTNode* temp = root->left ? root->left : root->right;

            if (temp == NULL)
            {
                temp = root;
                root = NULL;
            }
            else
                *root = *temp;
            free(temp);
        }
        else
        {
            BTNode* temp = minValueNode(root->right);
            root->key = temp->key;
            root->right = deleteNode(root->right, temp->key);
        }
    }

    if (root == NULL) return root;

    root->height = 1 + max(height(root->left), height(root->right));

    int balance = getBalance(root);

    //LL型
    if (balance > 1 && getBalance(root->left) >= 0) return ll_rotate(root);
    //LR型
    if (balance > 1 && getBalance(root->left) < 0) 
    {
        root->left = rr_rotate(root->left);
        return ll_rotate(root);
    }
    //RR型
    if (balance < -1 && getBalance(root->right) <= 0) return rr_rotate(root);
    //Rl型
    if (balance < -1 && getBalance(root->right) > 0) 
    {
        root->right = ll_rotate(root->right);
        return rr_rotate(root);
    }

    return root;
}

void preOrder(struct Node *root)
{
    if (root != NULL)
    {
        printf("(%d, %d) ", root->key, root->height);
        preOrder(root->left);
        preOrder(root->right);
    }
}

int main()
{
    BTNode *root = NULL;
    int list[9] = {9, 5, 10, 0, 6, 11, -1, 1, 2};
    int i;
    for(i = 0; i < 9; i ++)
    {
        root = insert(root, list[i]);
    }

    printf("前序遍历：\n");
    preOrder(root);

    /* The constructed AVL Tree would be
                     9
                    /  \
                   1    10
                 /  \     \
                0    5     11
               /    /  \
              -1   2    6
    */
    
    root = deleteNode(root, 10);
    /* The AVL Tree after deletion of 10
                       1
                     /   \
                    0     9
                  /     /  \
                -1     5     11
                     /  \
                    2    6
    */
    printf("\n");
    printf("前序遍历：\n");
    preOrder(root);
    printf("\n");
    return 0;
}

// gcc -o tree btree.c

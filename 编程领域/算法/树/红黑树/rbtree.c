// 源码来源[weewqrer 红黑树](http://blog.csdn.net/weewqrer/article/details/51866488)
#include <malloc.h>
#include <assert.h>

typedef enum ColorType {RED, BLACK} ColorType;
typedef struct rbt_t rbt_t;
typedef struct rbt_root_t rbt_root_t;

struct rbt_t{
    int value;
    rbt_t* left;
    rbt_t* right;
    rbt_t* p;
    ColorType color;
};

struct rbt_root_t{
    rbt_t* root;
    rbt_t* nil;
};

//函数声明
rbt_root_t* rbt_init(void);
static void rbt_handleReorient(rbt_root_t* T, rbt_t* x);
rbt_root_t* rbt_insert(rbt_root_t* T, int k);
rbt_root_t* rbt_delete(rbt_root_t* T, int k);

void rbt_transplant(rbt_root_t* T, rbt_t* u, rbt_t* v);

static rbt_t* rbt_rr_rotate( rbt_root_t* T, rbt_t* x);
static rbt_t* rbt_ll_rotate( rbt_root_t* T, rbt_t* x);

void rbt_inPrint(const rbt_root_t* T, rbt_t* t);
void rbt_prePrint(const rbt_root_t * T, rbt_t* t);
void rbt_print(const rbt_root_t* T);

static rbt_t* rbt_findMin(rbt_root_t * T, rbt_t* t);
static rbt_t* rbt_findMax(rbt_root_t * T, rbt_t* t);

static rbt_t* rbt_findMin(rbt_root_t * T, rbt_t* t){
    if(t == T->nil) return T->nil;

    while(t->left != T->nil) t = t->left;

    return t;
}
static rbt_t* rbt_findMax(rbt_root_t * T, rbt_t* t){
    if(t == T->nil) return T->nil;

    while(t->right != T->nil) t = t->right;

    return t;
}

/*
*@function 初始化, 创建根节点并为其分配空间, 同时创建共用的nil节点.
*/
rbt_root_t* rbt_init(void){
    rbt_root_t* T;

    T = (rbt_root_t*)malloc(sizeof(rbt_root_t));
    assert(NULL != T);

    // 创建nil节点并初始化, 该红黑树中所有nil节点都将指向这里.
    T->nil = (rbt_t*)malloc(sizeof(rbt_t));
    assert(NULL != T->nil);
    T->nil->color = BLACK;
    T->nil->left = T->nil->right = NULL;
    T->nil->p = NULL;

    T->root = T->nil;

    return T;
}

/*
*@function 内部函数 由rbt_insert调用
在第一种情况下, 进行颜色翻转; 在第二种情况下, 相当于对新插入的x点初始化
*/ 
void rbt_handleReorient(rbt_root_t* T, rbt_t* x){
    // 如果目标节点的父节点为黑, 则不用进行任何调整操作.
    if(x->p->color == BLACK) 
    {
        //无条件令根为黑色
        T->root->color = BLACK;
        return;
    }

    // 运行到这, 说明父节点为红, 那肯定不是根节点, 而且x的祖父节点肯定为黑.

    // 如果叔叔节点也为红, 则需要变换颜色, 否则只需要旋转.
    rbt_t* gp = x->p->p;
    if(gp->left->color == gp->right->color)
    {
        gp->color = RED;
        gp->left->color = gp->right->color = BLACK;

        if(gp == T->root)
        {
            gp->color = BLACK;
            return;
        }else if (gp->p->color == BLACK) {
            return;
        } else {
            return rbt_handleReorient(T, gp);
        }
    } else {
        rbt_t* subroot;

        // 父节点为红, 则祖父节点一定不为红, 这里把祖父节点也转为红色.
        // 进行旋转的时候, 一定会涉及到一次颜色转换(4种旋转都是)
        // 这里就是把作为 `A -> B -> C` 这一小段子树中的节点A的颜色转换一次.
        //此时x, x->p, x->p->p都为红, 另外还一个节点一定需要转换成黑色, 
        // 要看具体的旋转类型而定.
        x->p->p->color = RED; 
        // LR
        if(x->p->value < x->p->p->value && x->value > x->p->value) 
        {
            x->color = BLACK;
            subroot = rbt_rr_rotate(T,x->p);
            subroot = rbt_ll_rotate(T,x->p);
        }
        // LL
        else if(x->p->value < x->p->p->value && x->value < x->p->value) 
        {
            x->p->color = BLACK;
            subroot = rbt_ll_rotate(T,x->p->p);
        }
        // RR
        else if(x->p->value > x->p->p->value && x->value > x->p->value) 
        {
            x->p->color = BLACK;
            subroot = rbt_rr_rotate(T,x->p->p);
        }
        // RL
        else if(x->p->value > x->p->p->value && x->value < x->p->value) 
        {
            x->color = BLACK;
            subroot = rbt_ll_rotate(T,x->p);
            subroot = rbt_rr_rotate(T,x->p);
        }

        //无条件令根为黑色
        T->root->color = BLACK;
    }
}
/*
*@function brt_insert 插入
*1 新插入的结点默认为红, 如果为黑, 会破坏条件4(每个结点到null叶结点的每条路径有同样数目的黑色结点)
*2 如果新插入的结点的父节点为黑(不管叔叔节点), 那么插入完成. 如果父亲是红色的, 那么做一个旋转即可. (前提是叔叔是黑色的)
*3 我们这个插入要保证其叔叔是黑色的. 也就是在x下沉过程中, 不允许存在两个红色结点肩并肩. 
*/
rbt_root_t* rbt_insert(rbt_root_t* T, int val){
    rbt_t *x, *p;
    // 变量x指向目标节点要插入的位置.
    // 变量p一直指向遍历过程中当前节点的父节点.
    x = T->root;
    p = x;

    // 从根节点开始, 遍历红黑树, 寻找合适的插入点.
    // 根节点则会直接跳过这个while循环.
    // 令x下沉到叶子上, 而且保证一路上不会有同时为红色的兄弟
    while(x != T->nil){
        // 保证没有一对兄弟同时为红色, 交换父节点/叔叔节点与祖父节点的颜色.
        // 不过为什么在插入之前做这些事?
        // if(x->left->color == RED && x->right->color == RED)
        // {
        //     x->color = RED;
        //     x->left->color = x->right->color = BLACK; 
        //     rbt_handleReorient(T, x);
        // }

        p = x;
        if(val < x->value)
            x = x->left;
        else if(val>x->value)
            x = x->right;
        else{
            printf("\n%d已存在\n",val);
            return T;
        }
    }

    // 找到了新节点要插入的位置, 为其分配空间, 并进行初始化.
    // 注意: 这里x已经与T->root没有关系了, ta指向了新的地址.
    x = (rbt_t *)malloc(sizeof(rbt_t));
    assert(NULL != x);
    x->value = val;
    x->color = RED;
    x->left = x->right = T->nil;
    x->p = p;

    // 把新节点x插入树中, 作为p的子节点.
    // 如果当前树为空树, 则可以直接作为根节点. 此时根节点的p仍然是ta本身
    if(T->root == T->nil)
        T->root = x;
    else if(val < p->value)
        p->left = x;
    else
        p->right = x;

    // 插入完成, 开始调整.
    // 因为一路下来, 如果x的父亲是红色, 那么x的叔叔肯定不是红色了, 这个时候只需要做一下翻转即可. 
    rbt_handleReorient(T, x);

    return T;
}
void rbt_transplant(rbt_root_t* T, rbt_t* u, rbt_t* v){
    if(u->p == T->nil)
        T->root = v;
    else if(u == u->p->left)
        u->p->left =v;
    else
        u->p->right = v;
    v->p = u->p;
}
/*
*@brief rbt_delete 从树中删除 k
*/
rbt_root_t* rbt_delete(rbt_root_t* T, int k){
    assert(T != NULL);
    if(NULL == T->root) return T;

    //找到要被删除的叶子结点
    rbt_t * toDelete = T->root; 
    rbt_t * x;

    //找到值为k的结点
    while(toDelete != T->nil && toDelete->value != k){
        if(k<toDelete->value)
            toDelete = toDelete->left;
        else if(k>toDelete->value)
            toDelete = toDelete->right;
    }

    if(toDelete == T->nil){
        printf("\n%d 不存在\n",k);
        return T;
    }

    //如果两个孩子, 就找到右子树中最小的代替, alternative最多有一个右孩子
    if(toDelete->left != T->nil && toDelete->right != T->nil){
        rbt_t* alternative = rbt_findMin(T, toDelete->right);
        k = toDelete->value = alternative->value;
        toDelete = alternative;
    }

    if(toDelete->left == T->nil){
        x = toDelete->right;
        rbt_transplant(T,toDelete,toDelete->right);
    }else if(toDelete->right == T->nil){
        x = toDelete->left;
        rbt_transplant(T,toDelete,toDelete->left);
    }

    if(toDelete->color == BLACK){
        //x不是todelete, 而是用于代替x的那个
        //如果x颜色为红色的, 把x涂成黑色即可,  否则 从根到x处少了一个黑色结点, 导致不平衡
        while(x != T->root && x->color == BLACK){
            if(x == x->p->left){
                rbt_t* w = x->p->right;

                //情况1 x的兄弟是红色的, 通过
                if(RED == w->color){
                    w->color = BLACK;
                    w->p->color = RED;
                    rbt_rr_rotate(T,x->p);
                    w = x->p->right;
                }//处理完情况1之后, w.color== BLACK ,  情况就变成2 3 4 了

                //情况2 x的兄弟是黑色的, 并且其儿子都是黑色的. 
                if(w->left->color == BLACK && w->right->color == BLACK){
                    if(x->p->color == RED){
                        x->p->color = BLACK;
                        w->color = RED;
                        break;
                    }else{
                        w->color = RED;
                        x = x->p;//x.p左右是平衡的, 但是x.p处少了一个黑结点, 所以把x.p作为新的x继续循环
                        continue;
                    }
                }

                //情况3 w为黑色的, 左孩子为红色. (走到这一步, 说明w左右不同时为黑色. )
                if(w->right->color == BLACK){
                    w->left->color = BLACK;
                    w->color = RED;
                    rbt_ll_rotate(T,w);
                    w = x->p->right;
                }//处理完之后, 变成情况4

                //情况4 走到这一步说明w为黑色,  w的左孩子为黑色,  右孩子为红色. 

                w->color=x->p->color;
                x->p->color=BLACK;
                w->right->color=BLACK;
                rbt_rr_rotate(T,x->p);
                x = T->root;
            }else{
                rbt_t* w = x->p->left;
                //1
                if(w->color == RED){
                    w->color = BLACK;
                    x->p->color = RED;
                    rbt_ll_rotate(T,x->p);
                    w = x->p->left;
                }
                //2
                if(w->left->color==BLACK && w->right->color == BLACK){
                    if(x->p->color == RED){
                        x->p->color = BLACK;
                        w->color = RED;
                        break;
                    }else{
                        x->p->color = BLACK;
                        w->color = RED;
                        x = x->p;
                        continue;
                    }
                }

                //3
                if(w->left->color == BLACK){
                    w->color = RED;
                    w->right->color = BLACK;
                    w = x->p->left;
                }

                //4
                w->color=w->p->color;
                x->p->color = BLACK;
                w->left->color = BLACK;
                rbt_ll_rotate(T,x->p);
                x = T->root;
            }
        }
        x->color = BLACK;
    }

    //放心删除todelete 吧
    free(toDelete);

    return T;
}

/*
*@brief rbt_ll_rotate
*@param[in] 树根
*@param[in] 要进行旋转的结点
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
rbt_t* rbt_ll_rotate(rbt_root_t* T, rbt_t* x){
    rbt_t * tmp = x->left;
    x->left = tmp->right;

    if(T->nil != x->left)
        x->left->p = x;

    tmp->p = x->p;
    if(tmp->p == T->nil)
        T->root = tmp;
    else if(tmp->value < tmp->p->value)
        tmp->p->left= tmp;
    else
        tmp->p->right = tmp;

    tmp->right = x;
    x->p = tmp;

    return tmp;
}

/*
*@brief rbt_rr_rotate
*@param[in] T 树根
*@param[in] x 要进行旋转的结点
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
                            RR                         LL(这一步不重要)
                   4        =>             4           =>                  4
                 /   \                   /   \                           /   \
               3      5                 3     5                         2     5
              /                       /                               /   \
    node ->  1                       2  <- key                       1     3  
               \                   /                   
                2  <- key         1  <- node                  
*/
rbt_t* rbt_rr_rotate(rbt_root_t* T, rbt_t* node){
    rbt_t* tmp = node->right;

    node->right = tmp->left;
    if(node->right != T->nil) node->right->p = node;

    tmp->p = node->p;
    if(tmp->p == T->nil) // 这是哪种情况?
        T->root = tmp;
    else if(tmp->value < tmp->p->value)
        tmp->p->left = tmp;
    else
        tmp->p->right = tmp;

    tmp->left = node;
    node->p = tmp;

    return tmp;
}

// 前序遍历
void rbt_prePrint(const rbt_root_t* T, rbt_t* t){
    if(T->nil == t)return ;
    if(t->color == RED)
        printf("%3dR",t->value);
    else
        printf("%3dB",t->value);
    rbt_prePrint(T,t->left);
    rbt_prePrint(T,t->right);
}

// 中序遍历
void rbt_inPrint(const rbt_root_t* T, rbt_t* t){
    if(T->nil == t)return ;
    rbt_inPrint(T,t->left);
    if(t->color == RED)
        printf("%3dR",t->value);
    else
        printf("%3dB",t->value);
    rbt_inPrint(T,t->right);
}

//打印程序包括前序遍历和中序遍历两个, 因为它俩可以唯一确定一棵二叉树
void rbt_print(const rbt_root_t* T){
    assert(T!=NULL);
    printf("\n前序遍历 : ");
    rbt_prePrint(T,T->root);
    printf("\n中序遍历 : ");
    rbt_inPrint(T,T->root);
    printf("\n");
}

void main(){
    rbt_root_t* T = rbt_init();

    /************************************************************************/
    /* 1    测试插入
    /*
    /*
    /*输出  前序遍历 :   7B  2R  1B  5B  4R 11R  8B 14B 15R
    /*      中序遍历 :   1B  2R  4R  5B  7B  8B 11R 14B 15R
    /************************************************************************/

    // int list[9] = {11, 7, 1, 2, 8, 14, 15, 5, 4};
    // int length = 9;

    int list[20] = {12,1,9,2,0,11,7,19,4,15,18,5,14,13,10,16,6,3,8,17};
    int length = 20;

    // int list[19] = {12,1,9,2,0,11,7,19,4,15,18,5,14,13,10,16, 6, 3, 8};
    // int length = 19;
    int i;
    for(i = 0; i < length; i ++){
        T = rbt_insert(T, list[i]);
        printf("%d\n", list[i]);
        rbt_print(T);
    }

    // T = rbt_insert(T,4); //重复插入测试
    rbt_print(T);

    /************************************************************************/
    /* 2    测试删除
    /*    
    /*操作  连续删除4个元素 rbt_delete(T,8);rbt_delete(T,14);rbt_delete(T,7);rbt_delete(T,11);
    /*输出  前序遍历 :   2B  1B  5R  4B 15B
    /*      中序遍历 :   1B  2B  4B  5R 15B
    /************************************************************************/

    return;
    rbt_delete(T,8);
    rbt_delete(T,14);
    rbt_delete(T,7);
    rbt_delete(T,11);

    rbt_delete(T,8);//删除不存在的元素
    rbt_print(T);
}

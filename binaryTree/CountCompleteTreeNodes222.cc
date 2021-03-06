/**
 * 222. 完全二叉树的节点个数
 * 
 * 完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，
 * 并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。
输入: 
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6
 */
#include<math.h>
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
class Solution {
public:
    int countNodes(TreeNode* root) {
        TreeNode* l = root, *r = root;
        int lh = 0 , rh = 0;
        while(l){
            l = l->left;
            lh++;
        }
        while(r){
            r = r->right;
            rh++;
        }
        //满二叉树
        if(lh == rh) return (int)pow(2,lh) - 1;
        return 1 + countNodes(root->left) + countNodes(root->right);
    }
};
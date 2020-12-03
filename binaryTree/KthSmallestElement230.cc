/**
 * 230. 二叉搜索树中第K小的元素
 * 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
 * 说明：
 * 你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
 输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1
 * */

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution {
public:
    int rank = 0;
    int rlt = 0;
    int kthSmallest(TreeNode* root, int k) {
        traverse(root,k);
        return rlt;
    }
    //采用中序遍历的方式 进行排序
    void traverse(TreeNode* root, int k){
        if(root == nullptr) return;
        traverse(root->left, k);
        rank ++;
        if(k == rank){
            rlt = root->val;
            return;
        }
        traverse(root->right,k);
    }
};
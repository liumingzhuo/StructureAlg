/**
 * 226. Invert Binary Tree
 * 
     4
   /   \
  2     7
 / \   / \
1   3 6   9
output
     4
   /   \
  7     2
 / \   / \
9   6 3   1
 * */
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 };
class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        if(root == nullptr) return nullptr;
        //交换左右节点
        TreeNode *temp = root->left;
        root->left = root->right;
        root->right = temp;
        //递归遍历每一个叶子节点
        invertTree(root->left);
        invertTree(root->right);
    }
};
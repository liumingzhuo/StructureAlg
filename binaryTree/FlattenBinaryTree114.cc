/**
 * 114. Flatten Binary Tree to Linked List
 * 给定一个二叉树，原地将它展开为一个单链表。
 * 
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
    void flatten(TreeNode* root) {
        //base case
        if(root == nullptr)  return;
        flatten(root->left);
        flatten(root->right);

        TreeNode *left = root->left;
        TreeNode *right = root->right;

        //将左子树接到右子树
        root->right = left;
        root->left = nullptr;
        //将原来的右子树接到新的右子树上
        TreeNode* p = root;
        while(p ->right != nullptr)
            p = p->right;  
        p->right = right;
    }
};
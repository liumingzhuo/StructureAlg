/**
 * 450. 删除二叉搜索树中的节点
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。
返回二叉搜索树（有可能被更新）的根节点的引用。
一般来说，删除节点可分为两个步骤：
首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。
 * 
 */
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
    TreeNode* deleteNode(TreeNode* root, int key) {
        if(!root) return nullptr;
        if(root->val == key){
            //左节点为空
            if(!root->left) return root->right;
            //有节点为空
            if(!root->right) return root->left;
            //都不为空
            TreeNode* maxNode = root->right;
            while (maxNode->left) maxNode = maxNode->left;
            maxNode->left = root->left;
            return root->right;
 
        }else if(root->val > key){
            root->left = deleteNode(root->left, key);
        }else if(root->val < key){
            root->right = deleteNode(root->right,key);
        }
        return root;
    }
    TreeNode* getMax(TreeNode* node){
        while(node->left != nullptr){
            node = node->left;
        }
        return node;
    }
};

/**
 * 111. Minimum Depth of Binary Tree
 * Given a binary tree, find its minimum depth.
 * The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
 * Note: A leaf is a node with no children.
 * */
#include<queue>
using namespace std;
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
    int minDepth(TreeNode* root) {
        if (root == nullptr) return 0;
        queue<TreeNode*> nodeQueue;
        nodeQueue.push(root);
        int depth = 1;
        while (!nodeQueue.empty())
        {
            int n = nodeQueue.size();
            for(int i = 0; i< n; i++)
            {
                TreeNode *curNode = nodeQueue.front();
                nodeQueue.pop();
                if(curNode->left == nullptr && curNode->right == nullptr) return depth;
                if(curNode->left != nullptr) nodeQueue.push(curNode->left);
                if(curNode->right!= nullptr) nodeQueue.push(curNode->right);    
            }
            ++depth;
        }
        return depth;
    }

    /**
    int minDepth(TreeNode* root) {
        if(!root)
            return 0;
        int left=minDepth(root->left),right=minDepth(root->right);
        return (left && right) ? 1+min(left,right):1+left+right;
    }
     * */
};
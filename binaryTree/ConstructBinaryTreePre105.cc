/**
 * 105. 从前序与中序遍历序列构造二叉树
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。
 * */
#include<vector>
using namespace std;
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
class Solution {
public:
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        return helper(preorder,0,preorder.size()-1,
                    inorder,0,inorder.size());
    }

    TreeNode* helper(vector<int> &preorder,int preStart, int preEnd,
                    vector<int>& inorder, int inStart, int inEnd){
            if(preStart > preEnd) return NULL;
            int rootVal = preorder[preStart];
            int index = 0;
            for(int i = inStart; i <= inEnd; i++){
                if(inorder[i] == rootVal){
                    index = i;
                    break;
                }
            }
            // leftsize 用于preorder查找左子树
            int leftSize = index - inStart;
            TreeNode* root = new TreeNode(rootVal);
            //前序遍历第一个元素一定是根节点
            root->left = helper(preorder,preStart + 1, preStart + leftSize ,
                                inorder,inStart, index -1);
            root->right= helper(preorder,preStart+ leftSize+1,preEnd,
                                inorder,index + 1, inEnd);
            return root;
    }
};
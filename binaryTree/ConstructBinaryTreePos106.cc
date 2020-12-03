/**
 * 106. 从中序与后序遍历序列构造二叉树
 * 根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。
 * 
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
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        return  helper(inorder, 0, inorder.size() - 1,
                    postorder, 0, postorder.size() - 1);
    }
    TreeNode* helper(vector<int> &inorder,int inStart, int inEnd,
                    vector<int>& posorder, int posStart, int posEnd){

        if(inStart > inEnd) return nullptr;
        int rootVal = posorder[posEnd];
        int index = 0;
        for(int i = inStart; i <= inEnd; i++){
            if(inorder[i] == rootVal){
                index = i;
                break;
            }
        }
        int leftSize = index - inStart;
        TreeNode* root = new TreeNode(rootVal);
        //后序遍历 最后一位是根节点
        root->left = helper(inorder, inStart, index -1, posorder, posStart, posStart + leftSize - 1 );
        root->right = helper(inorder, index + 1, inEnd, posorder, posStart + leftSize, posEnd - 1);
        return root;
    }
};
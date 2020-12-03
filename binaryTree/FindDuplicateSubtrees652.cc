/**
 * 652. 寻找重复的子树
给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
两棵树重复是指它们具有相同的结构以及相同的结点值。
 * 
 * */
#include<string>
#include<iostream>
#include<unordered_map>
#include<vector>
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
    vector<TreeNode*> rlt;
    unordered_map<string, int> duplicateMap;
    vector<TreeNode*> findDuplicateSubtrees(TreeNode* root) {
        traverse(root);
        return rlt;
    }

    string traverse(TreeNode* root){
        if(root == nullptr) return "#";
        string left = traverse(root->left);
        string right = traverse(root->right);
        string subTree = left + "," + right + "," + to_string(root->val);
        auto flag = duplicateMap.find(subTree);

        if(flag == duplicateMap.end()){
            duplicateMap[subTree] = 0;
        }else if(flag -> second == 1){
            rlt.push_back(root);
        }
        duplicateMap[subTree] = ++ duplicateMap[subTree] ;
        return subTree;
    }
};
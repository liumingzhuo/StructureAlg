/**
 * 654. 最大二叉树
给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
二叉树的根是数组中的最大元素。
左子树是通过数组中最大值左边部分构造出的最大二叉树。
右子树是通过数组中最大值右边部分构造出的最大二叉树。
通过给定的数组构建最大二叉树，并且输出这个树的根节点。
 * 
 * */
#include<vector>
#include<iostream>
using namespace std;
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

class Solution {
public:
    TreeNode* constructMaximumBinaryTree(vector<int>& nums) {
        return helper(nums,0,nums.size()-1);
    }

    TreeNode* helper(vector<int> &nums, int lo, int hi){
        if(lo > hi) return nullptr;
        int maxValue = INT32_MIN;
        int index = 0;
        for(int i = 0; i < nums.size(); i++){
            if(maxValue < nums[i]){
                maxValue = nums[i];
                index = i;
            }
        }
        TreeNode* root =new TreeNode(maxValue);
        root->left = helper(nums, lo, index - 1);
        root->right = helper(nums,index + 1, hi);
        return root;
    }
};
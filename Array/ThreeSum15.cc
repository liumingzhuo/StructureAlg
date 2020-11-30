/**
 * 15. 3Sum
 * Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? 
 * Find all unique triplets in the array which gives the sum of zero.
 * Notice that the solution set must not contain duplicate triplets.
 * */
#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        return threeSumTarget(nums, 0);
    }
    vector<vector<int>> threeSumTarget(vector<int>& nums,int target) {
        sort(nums.begin(),nums.end());
        vector<vector<int>> res;
        int n = nums.size();
        for(int i = 0; i < n; i++){
            //注意target 应该减去当前nums[i]
            vector<vector<int>> tuples = twoSumTarget(nums, i+1, target - nums[i]);
            for(vector<int> &tuple:tuples){
                tuple.push_back(nums[i]);
                res.push_back(tuple);
            } 
            while(i < n-1 && nums[i] == nums[i+1]) i++;
        }
        return res;
    }
    vector<vector<int>> twoSumTarget(vector<int>& nums, int start, int target){
        //双指针  转为计算2数之合
        int lo = start, hi = nums.size()-1;
        vector<vector<int>> res;
        while (lo < hi){
            int sum = nums[lo] + nums[hi];
            int left = nums[lo] ,right = nums[hi];
            if(sum < target){
               while (lo < hi && nums[lo] == left) lo ++;
            }else if (sum > target){
                while(lo < hi && nums[hi] == right) hi--;
            }else{
                res.push_back(vector<int>({left,right}));
                while(lo < hi && nums[lo] == left) lo++;
                while(lo < hi && nums[hi] == right) hi--;
            }
        }
        return res;
        
    }
};
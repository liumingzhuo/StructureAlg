/**
 * 18.4Sum
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，
 * 判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？
 * 找出所有满足条件且不重复的四元组。
 * 注意：
 * 答案中不可以包含重复的四元组。
 * 
 * Example
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
 * */
#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        sort(nums.begin(),nums.end());
        return nSumTarget(nums, 4, 0, target);
    }   

    vector<vector<int>> nSumTarget(vector<int>& nums, int n, int start, int target){
            int sz = nums.size();
            vector<vector<int>> res;
            if(n < 2 || sz < n) return res;
            if(n == 2){
                int low = start, high = sz - 1;
                while (low <  high){
                    int sum = nums[low] + nums[high];
                    int left = nums[low], right = nums[high];
                    if(sum < target){
                        while(low < high && nums[low] == left) low++;
                    }else if(sum > target){
                        while(low < high && nums[high] == right) high--;
                    }else{
                        res.push_back({left,right});
                        while(low < high && nums[low] == left) low++;
                        while(low < high && nums[high]== right) high--;
                    }
                }
            }else{
                //从start 开始遍历
                for(int i = start; i < sz; i++){
                    vector<vector<int>> tuples = nSumTarget(nums, n - 1, i + 1, target - nums[i]);
                    for(vector<int> &tuple : tuples){
                        tuple.push_back(nums[i]);
                        res.push_back(tuple);
                    }
                    //跳过第一个数字重复情况
                    while( i < sz - 1 && nums[i] == nums[i+1]) i++;
                }
            }
            return res;
    }   
};
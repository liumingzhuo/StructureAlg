/**
 * 704. Binary Search
 * Given a sorted (in ascending order) integer array nums of n elements and a target value,
 *  write a function to search target in nums.
 *  If target exists, then return its index, otherwise return -1.
 * 
 * Example
 * Input: nums = [-1,0,3,5,9,12], target = 9
 * Output: 4
 * Explanation: 9 exists in nums and its index is 4
 * */
#include<vector>
#include<iostream>

using namespace std;
class Solution {
public:
    int search(vector<int>& nums, int target) {
        if (nums.size() == 0)return -1;
        int left = 0;
        int right = nums.size() - 1;
        while (left <= right){
            int mid = left + (right - left)/2;
            if(nums[mid] < target){
                left = mid + 1;
            }else if (nums[mid] > target)
            {
                right = mid - 1;
            }else if (nums[mid] ==  target){
                return mid;
            }
        }
        return -1;
    }
};
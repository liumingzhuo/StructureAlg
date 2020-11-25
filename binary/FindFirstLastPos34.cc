/**
 * 34. Find First and Last Position of Element in Sorted Array
 * Given an array of integers nums sorted in ascending order, 
 * find the starting and ending position of a given target value.
 * If target is not found in the array, return [-1, -1].
 * Follow up: Could you write an algorithm with O(log n) runtime complexity?
 * 
 * Example
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 * */
#include<vector>
#include<iostream>
using namespace std;
class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        vector<int> rlt(2);
        int left = findPos(nums, target, true);
        int right =  findPos(nums, target, false);
        rlt[0] = left;
        rlt[1] =right;
        return rlt;
    }
    int findPos(vector<int>&nums,int target,bool isLeft){
        int left = 0;   
        int right = nums.size() - 1;
        while (left <= right){
            int mid = left + (right - left)/2;
            if (nums[mid] < target){
               left = mid + 1;
            }else if (nums[mid] > target){
                right = mid - 1;
            }else if (nums[mid] == target){
                if(isLeft){
                    right = mid - 1;
                } 
                else{
                    left = mid + 1;
                }
            }
        }
        if(isLeft){
            if (left >= nums.size() || nums[left]!= target) 
                return -1;
            return left;
        }else
        {
            if (right < 0 || nums[right] != target) 
                return -1;
            return right;
        }

        
    }
};

/**
 * 46. Permutations
 * Given an array nums of distinct integers, 
 * return all the possible permutations. You can return the answer in any order.
 * 
 * Example 
 * Input: nums = [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 * */
#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;
class Solution {
public:
    vector<vector<int> >res;
    /**
    vector<vector<int>> permute(vector<int>& nums) {
        vector<int> track;
        vector<bool> isContains(nums.size(),false);
        backtrack(nums,track,isContains);
        return res;
    }
    void backtrack(vector<int> &nums, vector<int> &track,vector<bool> &isContains){
        if (nums.size() == track.size()){
            res.push_back(track);
            return;
        }
        for(int i=0;i<nums.size();i++){
            //剪枝
            if(isContains[i]) continue;
            track.push_back(nums[i]);
            isContains[i] = true;
            backtrack(nums,track,isContains);
            track.pop_back();
            isContains[i] = false;
        }
    }
    **/

    vector<vector<int>> permute(vector<int>& nums) {
		func(nums, 0, nums.size()-1);
		return res;
	}
    int func(vector<int> &nums,int start,int end){
        if (start == end){
            res.push_back(nums);
            return;
        }
        for (int i = start; i <= end; i++){
           swap(nums[i],nums[start]);
           func(nums,start+1,end);
           swap(nums[i],nums[start]);
        }  
    }
};

int main(){
    Solution s = Solution();
    int a[3]={1,2,3};
    vector<int> vec(a,a+3);
    vector<vector<int>> rlt = s.permute(vec);
    for(auto it = rlt.begin();it<rlt.end();it++){
        for (auto itt = (*it).begin();itt < (*it).end(); itt++){
                cout<<*itt<<endl;
        }   
    }
}
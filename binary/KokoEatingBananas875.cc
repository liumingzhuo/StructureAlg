/**
 * 875. 爱吃香蕉的珂珂
珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。

珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，
她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。  

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。
输入: piles = [3,6,7,11], H = 8
输出: 4
 */ 

#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    int minEatingSpeed(vector<int>& piles, int H) {
        int max_num = *max_element(piles.begin(), piles.end());
        int left = 1, right = max_num + 1;
        while (left < right){
            int mid = left + (right - left)/2;
            if(finish(piles,mid,H)){
                right = mid;
            }else{
                left = mid + 1;
            }
        }
        return left; 
        
    }
    bool finish(vector<int>& piles, int speed, int H){
        int time = 0;
        for(int n :piles){
            time += (n/speed) + (n%speed > 0 ? 1 : 0);
        }
        return time <= H;
    }
};
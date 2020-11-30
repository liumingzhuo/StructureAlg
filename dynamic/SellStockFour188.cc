/**
 * 188. Best Time to Buy and Sell Stock IV
 * You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
 * Design an algorithm to find the maximum profit. You may complete at most k transactions.
 * Notice that you may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
 * 
 * Example
 * Input: k = 2, prices = [2,4,1]
 * Output: 2
 * Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
 * */
#include<vector>
#include<algorithm>
#include<iostream>
using namespace std;
class Solution {
public:
    int maxProfit(int k, vector<int>& prices) {
       int n = prices.size();
        //无限买卖
        if(k > n/2){
            int res = 0;
            for(int i = 1; i < n ;i++){
                if(prices[i]>prices[i-1]){
                    res += prices[i] - prices[i-1];
                }
            }
            return res;
        }
        //动态规划
        vector<int> dp_i_0(k+1, 0);
        vector<int> dp_i_1(k+1, INT32_MIN);
        for (auto price :prices) {
            for (int j = 1; j < k+1; j++){
                dp_i_0[j] = max(dp_i_0[j], dp_i_1[j] + price);
                dp_i_1[j] = max(dp_i_1[j], dp_i_0[j-1]-price);
            }
        }
        return dp_i_0[k];
    }
};

int main(){
    int input[6]= {3,2,6,5,0,3};
    vector<int> aaa (input,input+6);
    Solution s = Solution();
    int rlt = s.maxProfit(2,aaa);
    cout<<rlt<<endl;
}
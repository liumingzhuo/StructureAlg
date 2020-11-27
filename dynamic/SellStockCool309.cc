/**
 * 309. Best Time to Buy and Sell Stock with Cooldown
 * 
 * Say you have an array for which the ith element is the price of a given stock on day i.
 * Design an algorithm to find the maximum profit.
 * You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:
 * You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
 * After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
 * 
 * Input: [1,2,3,0,2]
 * Output: 3 
 * Explanation: transactions = [buy, sell, cooldown, buy, sell]
 * */
#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        //dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + price[i])
        //dp[i][k][1] = max(dp[i-1][k][1], dp[i-2][k-1][0] - price[i])
        int dp_i_0 = 0, dp_i_1 = INT16_MIN;
        //冻结期
        int dp_cold_down = 0;
        for(auto price : prices){
            int flag = dp_i_0;
            dp_i_0 = max(dp_i_0, dp_i_1 + price);
            dp_i_1 = max(dp_i_1, dp_cold_down - price);
            dp_cold_down = flag;
        }
        return dp_i_0;
    }
};
/**
 * 123. Best Time to Buy and Sell Stock III
 * Say you have an array for which the ith element is the price of a given stock on day i.
 * Design an algorithm to find the maximum profit. You may complete at most two transactions.
 * Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
 * 
 * */
#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        //dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + price[i])
        //dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - price[i])
        //dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + price[i]);
        //dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - price[i]);
        int dp_i_20 = 0, dp_i_21 = INT16_MIN;
        int dp_i_10 = 0, dp_i_11 = INT16_MIN;
        for(auto price : prices){
            dp_i_20 = max(dp_i_20, dp_i_21 + price);
            dp_i_21 = max(dp_i_21, dp_i_10 - price);
            dp_i_10 = max(dp_i_10, dp_i_11 + price);
            dp_i_11 = max(dp_i_11, - price);
        }
        return dp_i_20;
    }
};
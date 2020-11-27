/**
 * 714. Best Time to Buy and Sell Stock with Transaction Fee
 * Your are given an array of integers prices, for which the i-th element is the price of a given stock on day i; and a non-negative integer fee representing a transaction fee.
 * You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction. You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)
 * Return the maximum profit you can make.
 * 
 * Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
 * Output: 8
 * Explanation: The maximum profit can be achieved by:
 * Buying at prices[0] = 1
 * Selling at prices[3] = 8
 * Buying at prices[4] = 4
 * Selling at prices[5] = 9
 * The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
 * */

#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    int maxProfit(vector<int>& prices, int fee) {
        //dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + price[i])
        //dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - price[i] -fee)
        int dp_i_0 = 0, dp_i_1 = INT32_MIN; //注意是32位min
        for(int price : prices){
            int flag = dp_i_0;
            dp_i_0 = max(dp_i_0, dp_i_1 + price);
            dp_i_1 = max(dp_i_1, flag - price - fee);
        }
        return dp_i_0;
    }
};
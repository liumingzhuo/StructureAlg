/**
 * 122. Best Time to Buy and Sell Stock II
 * Say you have an array prices for which the ith element is the price of a given stock on day i.
 * Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
 * Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
 * Input: [7,1,5,3,6,4]
 * Output: 7
 * Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
 **/
#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        //dp[i][k][0] = max(dp[i-1][k][0],dp[i-1][k-1][1] + price[i]);
        //dp[i][k][1] = max(dp[i-1][k][1],dp[i-1][k-1][0] - price[i])
        //因为 k为无穷大  那么k-1 = k
        //所以 dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k][0] - price[i])
        int dp_i_0 = 0;
        int dp_i_1 = INT16_MIN;
        for(int price:prices){
            int flag = dp_i_0;
            dp_i_0 = max(dp_i_0, dp_i_1 + price);
            dp_i_1 = max(dp_i_1, flag-price);
        }
        return dp_i_0;
    }
};
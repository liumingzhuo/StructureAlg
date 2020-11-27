/**
 * 121. Best Time to Buy and Sell Stock
 * Say you have an array for which the ith element is the price of a given stock on day i.
 * If you were only permitted to complete at most one transaction 
 * (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
 * Note that you cannot sell a stock before you buy one.
 * 
 * Example:
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 * */
#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        // int n = prices.size();
        // int dp[n][2];
        // for (int i = 0; i < n; i++){
        //     if (i - 1 ==-1)
        //     {
        //        dp[i][0] = 0;
        //        //解释
        //        //dp[i][0] = max(dp[-1][0],dp[-1][1] + prices[i]);
        //        //           = max(0, 负无穷)
        //        dp[i][1] = -prices[i];
        //        //解释
        //        //dp[i][1] = max(dp[-1][1],dp[-1][0]-prices[i]);
        //        //         = max(负无穷,0- prices[i])
        //        continue;
        //     }
        //     dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i]);
        //     dp[i][1] = max(dp[i-1][1], -prices[i]);
        // }
        //return dp[n-1][0];

        //方式2优化 空间复杂度O(1)
        int dp_i_0 = 0,dp_i_1 = INT16_MIN;
        for(int price :prices){
            dp_i_0 = max(dp_i_0, dp_i_1 +price);
            dp_i_1 = max(dp_i_1, -price);
        }
        return dp_i_0;
    }
};
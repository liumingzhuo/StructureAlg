/**
 * 322. Coin Change
 * You are given coins of different denominations and a total amount of money amount.
 *  Write a function to compute the fewest number of coins that you need to make up that amount.
 *  If that amount of money cannot be made up by any combination of the coins, return -1.
 * You may assume that you have an infinite number of each kind of coin.
 * 
 * Example
 * Input: coins = [1,2,5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 * */
#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;
class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        vector<int> dp (amount+1, amount+1);
        dp[0] = 0;
        for (int i = 0; i < dp.size(); i++)
        {
           for(int coin : coins){
               if (i - coin < 0) continue;
               dp[i] = min(dp[i],1 + dp[i - coin]);
           }
        }
        return (dp[amount] == amount+1) ? -1 : dp[amount];
    }
};

int main(){
    Solution s = Solution();
    int coins_arr[3] ={1,2,5};
    vector<int> coins(coins_arr,coins_arr+3);
    cout<< s.coinChange(coins,11)<<endl;
}
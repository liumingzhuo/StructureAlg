/**
 * 509. Fibonacci Number
 * The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, 
 * such that each number is the sum of the two preceding ones, 
 * starting from 0 and 1. That is,
 * F(0) = 0,   F(1) = 1
 * F(N) = F(N - 1) + F(N - 2), for N > 1.
 * example Input: 2
 * Output: 1
 * Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
 * 
 * */
#include<vector>
#include<iostream>
using namespace std;
class Solution {
public:
    // int fib(int N) {
    //     if (N < 1) return 0;
    //     //创建一个备忘录，存储计算过的值，减少时间复杂度
    //     vector<int>temp(N+1,0);
    //     return decorator(temp,N);
    // }
    // int decorator(vector<int> &temp,int n){
    //     if (n ==1 || n==2) return 1;
    //     if (temp[n] != 0)  return temp[n];
    //     temp[n] = decorator(temp,n-1) + decorator(temp,n-2);
    //     return temp[n];
    // }   
    
    int fib(int N){ //迭代解法
        if (N ==1 || N ==2) return 1;
        int prev = 1,cur = 1;
        for(int i =3; i<=N ;i++){
            int sum = prev + cur;
            prev = cur;
            cur = sum;
        }
        return cur;
    }
};

int main(){
    Solution s = Solution();
    int fibRlt = s.fib(3);
    cout <<fibRlt <<endl;

}
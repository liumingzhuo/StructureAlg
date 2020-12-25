/**
 * 3. Longest Substring Without Repeating Characters
 * Given a string s, find the length of the longest substring without repeating characters.
 * 
 * Example
 * Input: s = "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 * */
#include<string>
#include<algorithm>
#include<unordered_map>
using namespace std;
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        unordered_map<char, int> windows;
        int left = 0, right = 0;
        int res = 0;
        while (right < s.size()){
            char c = s[right];
            right ++;
            windows[c]++;
            //窗口元素大于1 说明有重复 缩小窗口
            while (windows[c] > 1){
                char d = s[left];
                left ++;
                windows[d]--;
            }
            res = max(res, right - left);
        }
        return res;
    }
};
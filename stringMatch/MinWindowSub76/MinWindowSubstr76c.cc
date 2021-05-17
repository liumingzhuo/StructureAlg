/**
 * 76. Minimum Window Substring
 * Given two strings s and t, return the minimum window in s which will contain all the characters in t. 
 * If there is no such window in s that covers all characters in t, return the empty string "".
 * Note that If there is such a window, it is guaranteed that there will always be only one unique minimum window in s.
 * Input: s = "ADOBECODEBANC", t = "ABC"
 * Output: "BANC"
 * */
#include<string>
#include<unordered_map>
using namespace std;
class Solution {
public:
    string minWindow(string s, string t) {
        unordered_map<char, int > needs, windows;
        for(char c: t) needs[c]++;
        int left = 0,right = 0;
        //窗口中匹配到的字符的长度
        int valid = 0;
        int start = 0,len = INT16_MAX;
        while (right < s.size()){
            char c = s[right];
            //窗口右移
            right ++;
            //匹配到字符就加入到窗口
            if (needs.count(c)){
                windows[c]++;
                if (windows[c] == needs[c]){
                    valid ++;
              }
            }
            //当匹配长度和目标字符相等时
            while (valid == needs.size()){
                if (right - left < len){
                   start = left;
                   len = right - left;
                }
                char d = s[left];
                left ++;
                if(needs.count(d)){
                    if(windows[d] == needs[d])
                        valid --;
                    windows[d]--;
                }
            }
        }
        return len == INT16_MAX ? "" : s.substr(start,len);
        
    }
};
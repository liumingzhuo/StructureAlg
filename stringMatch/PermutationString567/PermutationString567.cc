/**
 * 567. Permutation in String
 * Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1.
 * In other words, one of the first string's permutations is the substring of the second string.
 * 
 * example 
 * Input: s1 = "ab" s2 = "eidbaooo"
 * Output: True
 * Explanation: s2 contains one permutation of s1 ("ba").
 * */
#include<string>
#include<iostream>
#include<unordered_map>
using namespace std;
class Solution {
public:
    bool checkInclusion(string s1, string s2) {
        unordered_map<char, int> needs,windows;
        for(char c:s1) needs[c]++;
        int left = 0, right = 0;
        int valid = 0;
        while (right < s2.size()){
            char c = s2[right];
            right ++;
            if (needs.count(c)){
                windows[c]++;
                if (windows[c] ==  needs[c])
                    valid++; 
            }
            //当窗口大小大于 s1.size 时 缩小窗口
            while (right - left >= s1.size()){
                //合法的排列
                if (valid == needs.size()) return true;
                char d = s2[left];
                left ++;
                if(needs.count(d)){
                    if(windows[d] == needs[d])
                        valid --;
                    windows[d]--;
                }
            }
        }
        return false;        
    }
};
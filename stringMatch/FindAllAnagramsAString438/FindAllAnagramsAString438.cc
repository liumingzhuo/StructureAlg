/**
 * 438. Find All Anagrams in a String
 * Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.
 * Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.
 * The order of output does not matter.
 * 
 * example
 * Input:
 * s: "cbaebabacd" p: "abc"
 * Output:
 * [0, 6]
 * Explanation:
 * The substring with start index = 0 is "cba", which is an anagram of "abc".
 * The substring with start index = 6 is "bac", which is an anagram of "abc".
 * */
#include<vector>
#include<string>
#include<unordered_map>
using namespace std;
class Solution {
public:
    vector<int> findAnagrams(string s, string p) {
        unordered_map<char, int> needs,windows;
        for(char c:p) needs[c]++;
        int left = 0,right = 0;
        int valid = 0;
        vector<int> rlt;
        while (right < s.size()){
            char c = s[right];
            right++;
            // 进行窗口内数据更新
            if (needs.count(c)){
                windows[c]++;
                if(needs[c] == windows[c])
                    valid++;
            }
            // 判断左窗口是否要收缩
            while (right - left >= p.size()){
                if(valid == needs.size())
                    rlt.push_back(left);
                char d = s[left];
                left ++;
                if(needs.count(d)){
                    if(needs[d] == windows[d])
                        valid--;
                    windows[d]--;
                }
            }
        }
        return rlt;
        
    }
};
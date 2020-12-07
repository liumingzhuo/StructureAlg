/**
 * 1002 Find Common Characters
  Given an array A of strings made only from lowercase letters, 
return a list of all characters that show up in all strings within the list (including duplicates). 
For example, 
if a character occurs 3 times in all strings but not 4 times, 
you need to include that character three times in the final answer.
You may return the answer in any order.
Example：
Input: ["bella","label","roller"]
Output: ["e","l","l"]
Input: ["cool","lock","cook"]
Output: ["c","o"]
 **/
#include<vector>
#include<iostream>
#include<string>
#include<algorithm>
using namespace std;
class Solution {
public:
    vector<string> commonChars(vector<string>& A) {
        string loopstr= A[0];
        int res [26] = {0};
        for(int i = 0;i<loopstr.length();i++ ) res[loopstr[i]-'a']++;
        for(auto it =A.begin()+1; it!=A.end();it++){
            int temp[26] = {0};
            for(int j = 0;j < (*it).length();j++) temp[(*it)[j]-'a']++;
            for(int j = 0;j < 26; j++){
                if(res[j]==0 && temp[j]==0)continue;
                res[j] = min(res[j],temp[j]);
            }
        }
        vector<string> rlt;
        for(int j = 0; j <26; j++){
            if(res[j] == 0) continue;
            for(int k = 0; k < res[j]; k++){
                string s = "";
                s= j+'a';
                rlt.push_back(s);
            }
        }
        return rlt;
    }
};

int main(){
    Solution sol = Solution();
    vector<string> A= {"bella","label","roller"};
    vector<string> rlt = sol.commonChars(A);
    for(string r :rlt){
        cout<<r<<endl;
    }
}

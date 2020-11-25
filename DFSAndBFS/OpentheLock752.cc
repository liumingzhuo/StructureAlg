/**
 * 752. Open the Lock 
 * You have a lock in front of you with 4 circular wheels. 
 * Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. 
 * The wheels can rotate freely and wrap around: 
 * for example we can turn '9' to be '0', or '0' to be '9'. 
 * Each move consists of turning one wheel one slot.
 * The lock initially starts at '0000', a string representing the state of the 4 wheels.
 * You are given a list of deadends dead ends, meaning if the lock displays any of these codes, 
 * the wheels of the lock will stop turning and you will be unable to open it.
 * Given a target representing the value of the wheels that will unlock the lock, 
 * return the minimum total number of turns required to open the lock, or -1 if it is impossible.
 * 
 * Example 
 * Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
 * Output: 6
 * Explanation:
 * A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
 * Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
 * because the wheels of the lock become stuck after the display becomes the dead end "0102".
 **/
#include<vector>
#include<string>
#include<cstring>
#include<iostream>
#include<queue>
#include<unordered_map>
#include<unordered_set>
#include<algorithm>
using namespace std;
class Solution {
public:
    int openLock(vector<string>& deadends, string target) {
        // unordered_map<string,int> visited;
        // for (auto i:deadends)
        //    visited[i] = 1;
        // if (visited["0000"] ==1) return -1;
        // queue<string> q;
        // q.push("0000"); 
        // int step = 0;

        // while (!q.empty())
        // {
        //     int sz = q.size();
        //     for (int i = 0; i< sz; i++){
        //         string cur = q.front();
        //         q.pop();
        //         if (cur == target) return step;
        //         if (visited[cur] == 1) continue;
        //         visited[cur] =1;
        //         for(int j = 0; j< 4; j++){
        //             string up = cur;
        //             string down = cur;
        //             up[j] = up[j] == '9'?'0':up[j]+1;
        //             down[j] = down[j] =='0'?'9':down[j]-1;
        //             if (visited[up]!= 1) q.push(up);
        //             if (visited[down]!=1)q.push(down);
        //         }
        //     }
        //     step ++;
        // }
        // return -1;



        /**
         * 双向BFS
         * */
        unordered_map<string, int> visited; //visited既是死亡列表也是走过的路径
        for(auto it:deadends)
            visited[it] = 1;
        if (visited["0000"] == 1) return -1;
        unordered_set<string> q1;
        unordered_set<string> q2;
        q1.insert("0000");
        q2.insert(target);
        int step = 0;
        while (!q1.empty() && !q2.empty())
        {
            //每次外层循环取最小的集合
            if (q1.size() > q2.size())
                q1.swap(q2);
            unordered_set<string> temp;
            for(auto iter = q1.begin(); iter != q1.end(); ++iter){
                if (q2.find(*iter) != q2.end()) return step;
                if (visited[*iter] == 1) continue;
                visited[*iter] = 1;
                for(int j = 0; j < 4; j++){
                    string up = *iter;
                    string down = *iter;
                    up[j] = up[j] == '9'?'0':up[j]+1;
                    down[j] = down[j] == '0'?'9':down[j]-1;
                    if (visited[up] != 1) temp.insert(up);
                    if (visited[down]!= 1) temp.insert(down);
                }
            }
            step ++;
            //交换q1 q2元素  当前q1下一次循环等于遍历q2
            q1 =q2;
            q2=temp;
        }
        
        return -1;
    }
};
/**
 *990.等式方程的可满足性 (并查集的思想解决)
给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，
并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。 
 */
#include<vector>
#include<string>
using namespace std;
class Solution {
public:
    int parent[];
    int capacity[];
    bool equationsPossible(vector<string>& equations) {
        for(int i = 0;i < 26; i++){
            parent[i] = i;
            capacity[i] = 1;
        }
        for(string eq : equations){
            if(eq.at(1) == '='){
                char x = eq.at(0);
                char y = eq.at(3);
                unionE(x - 'a',y - 'a');
            }
        }
        for(string eq : equations){
            if(eq.at(1) == '!'){
                char x = eq.at(0);
                char y = eq.at(3);
                if(connected(x - 'a', y - 'a')) return false;
            }
        }
        return true;
    }
    //节点联通
    void unionE(int p, int q){
        int rootP = find(p);
        int rootQ = find(q);
        if(rootP == rootQ) return;
        if(capacity[rootP] > capacity[rootQ]){
            parent[rootQ] = rootP; 
            capacity[rootP] += capacity[rootQ];
        }else{
            parent[rootP] = rootQ;
            capacity[rootQ] += capacity[rootP];
        }
    }
    bool connected(int p , int q){
        int rootP = find(p);
        int rootQ = find(q);
        return rootP == rootQ;
    }
    //返回节点x的根节点
    int find(int x){
        while(parent[x] != x){
            parent[x] = parent[parent[x]];
            x = parent[x];
        }
        return x;
    }
};
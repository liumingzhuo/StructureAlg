/**
 * 986. Interval List Intersections
 * 给定两个由一些 闭区间 组成的列表，每个区间列表都是成对不相交的，并且已经排序。
 * 返回这两个区间列表的交集。
 * （形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b。
 * 两个闭区间的交集是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3]。）
 * 
 * 
 * Example
 * 输入：A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
 * 输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
 * */
#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    vector<vector<int>> intervalIntersection(vector<vector<int>>& A, vector<vector<int>>& B) {
        vector<vector<int>> rlt;
        int i = 0,j=0;
        while (i < A.size() && j < B.size())
        {
            int a1 = A[i][0];
            int a2 = A[i][1];
            int b1 = B[j][0];
            int b2 = B[j][1];
            if(a1 <= b2 && b1 <= a2){ // 两个向量不香蕉条件  a2 < b1  || a1 > b2
                int temp[] ={max(a1,b1),min(a2,b2)};
                rlt.push_back(vector<int>(temp,temp+2));
            }
            if(b2 < a2) j++;
            else i++;
        }
        return rlt;
        
    }
};
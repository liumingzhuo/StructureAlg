/**
 * 56. Merge Intervals
 * Given an array of intervals where intervals[i] = [starti, endi], 
 * merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
 * 
 * Example
 * Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
 * */
#include<vector>
#include<bitset>
#include<algorithm>
using namespace std;
class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        vector<vector<int> >res;
        if(intervals.size() == 0 )return res;
        sort(intervals.begin(), intervals.end(), [](vector<int> a, vector<int> b){
            return a[0] < b[0];
        });
       
        res.push_back(intervals[0]);
        for (int i = 1; i < intervals.size(); i++)
        {
            vector<int> curr = intervals[i];
            vector<int> last = res[res.size()-1];
            if(curr[0] <= last[1]){
                last[1] = max(last[1], curr[1]);
                res[res.size()-1] = last;// 更新数据
            }else{
                res.push_back(curr);
            }
        }
        return res;
    }
};


// class Solution {
//     public int[][] merge(int[][] intervals) {
//         BitSet bitSet = new BitSet();
//         int max = 0;
//         for (int[] interval : intervals) {
//             int temp = interval[1] * 2 + 1;
//             bitSet.set(interval[0] * 2, temp, true);
//             max = temp >= max ? temp : max;
//         }

//         int index = 0, count = 0;
//         while (index < max) {
//             int start = bitSet.nextSetBit(index);
//             int end = bitSet.nextClearBit(start);

//             int[] item = {start / 2, (end - 1) / 2};
//             intervals[count++] = item;

//             index = end;
//         }
//         int[][] ret = new int[count][2];
//         for (int i = 0; i < count; i++) {
//             ret[i] = intervals[i];
//         }

//         return ret;
//     }
// }



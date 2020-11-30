/**
 * 1288. Remove Covered Intervals
 * Given a list of intervals, remove all intervals that are covered by another interval in the list.
 * Interval [a,b) is covered by interval [c,d) if and only if c <= a and b <= d.
 * After doing so, return the number of remaining intervals.
 * 
 * Example
 * Input: intervals = [[1,4],[3,6],[2,8]]
 * Output: 2
 * Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.
 * 
 * */
#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    int removeCoveredIntervals(vector<vector<int>>& intervals) {
        sort(intervals.begin(),intervals.end(),[](vector<int> a, vector<int> b){
            if(a[0] == b[0]) return a[1] > b[1];
            return a[0] < b[0];
        });
        int left = intervals[0][0];
        int right= intervals[0][1];
        int res = 0;
        for(int i = 1; i < intervals.size(); i++){
            vector<int> intev = intervals[i];
            // 在区间内
            if(left <= intev[0] && right >= intev[1]){
                res++;
            }
            //区间部分覆盖
            if(right >= intev[0] && right <= intev[1]){
                right = intev[1];
            }
            //区间不覆盖
            if(right < intev[0]){
                left = intev[0];
                right = intev[1];
            } 
        }
        return intervals.size() - res;
    }

};
/**
 * 621. 任务调度器
 * 给你一个用字符数组 tasks 表示的 CPU 需要执行的任务列表。其中每个字母表示一种不同种类的任务。
 * 任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。在任何一个单位时间，CPU 可以完成一个任务，或者处于待命状态。

然而，两个 相同种类 的任务之间必须有长度为整数 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

你需要计算完成所有任务所需要的 最短时间 。
 */
#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    int leastInterval(vector<char>& tasks, int n) {
        // 存在空闲时间 = (最大任务数 -1)*（冷却期+单个任务时间）+ 最后一个任务时间
        // 不存在空闲时间 = tasks.size
        int words[26] = {0};
        int length = tasks.size();
        //lastcount至少执行一次
        int lastCount = 1;
        for(auto task : tasks) ++words[task-'A'];
        //排序  words降序排列，准备第一个元素大小的桶的个数
        sort(words,words+26,[](int &x, int &y) -> int{ return x > y;});
        while(lastCount < 26 && words[lastCount] == words[0]) lastCount++;
        return max(length,(words[0]-1)*(n+1) + lastCount);
    }   
};
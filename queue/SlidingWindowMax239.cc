/**
 *239. 滑动窗口最大值
 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。

使用单调队列来实现滑动窗口问题
 */
#include<vector>
#include<list>
using namespace std;
class SingleQueue{
private:
    list<int> glist;
public:
    void push(int num){
        /** 删掉链表中 所有小于num 的值  即 链表是递减排序的    */
        while(!glist.empty() && glist.back() < num){
            glist.pop_back();
        }
        glist.push_back(num);
    }
    void pop(int num){
        if(num == glist.front()){
            glist.pop_front();
        }
    }
    int max(){
        return glist.front();
    }
}; 
class Solution {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        SingleQueue* q = new SingleQueue();
        vector<int> res;
        int index_k = k - 1;
        for(int i = 0; i < nums.size(); i++){
            if(i < index_k){
                q->push(nums[i]);
            }else{
                q->push(nums[i]);
                res.push_back(q->max());
                q->pop(nums[i-index_k]);
            }
        }
        delete q;
        return res;
    }
};
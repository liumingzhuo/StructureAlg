/**
 * 341. 扁平化嵌套列表迭代器
 * 给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。
列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。
输入: [[1,1],2,[1,1]]
输出: [1,1,2,1,1]
解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
 */

#include<vector>
#include<stack>
using namespace std;
class NestedInteger {
  public:
    // Return true if this NestedInteger holds a single integer, rather than a nested list.
    bool isInteger() const;
    // Return the single integer that this NestedInteger holds, if it holds a single integer
    // The result is undefined if this NestedInteger holds a nested list
    int getInteger() const;
    // Return the nested list that this NestedInteger holds, if it holds a nested list
    // The result is undefined if this NestedInteger holds a single integer
    const vector<NestedInteger> &getList() const;
};


class NestedIterator {
public:
    stack<NestedInteger *> stacks;
    NestedIterator(vector<NestedInteger> &nestedList) {
        for(auto it = nestedList.rbegin();it!=nestedList.rend();it++){
            stacks.push(&(*it));
        }
    }
    
    int next() {
        int rlt = stacks.top()->getInteger();
        stacks.pop();
        return rlt;
    }
    
    bool hasNext() {
        while (!stacks.empty() && !stacks.top()->isInteger())
        {
            vector<NestedInteger> &vec = stacks.top()->getList();
            stacks.pop();
            for(int i = vec.size() -1;i >= 0;i--)
                stacks.push(&vec[i]);
        }
        return !stacks.empty();
    }
};

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i(nestedList);
 * while (i.hasNext()) cout << i.next();
 */
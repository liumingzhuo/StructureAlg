#include <iostream>
#include <string>
#include <stack>
using namespace std;

// 20
class Solution
{
public:
    bool isValid(string s)
    {
        int len = s.size();
        stack<char> stk;
        if (len == 0)
            return true;
        if (s[0] == ')' || s[0] == ']' || s[0] == '}' || len % 2 != 0)
            return false;
        int a = 0;
        for (int i = 0; i < len; i++)
        {
            if (s[i] == '(' || s[i] == '[' || s[i] == '{')
            {
                stk.push(s[i]);
                continue;
            }
            if (stk.size() > 0)
            {
                char top = stk.top();
                if (top == '(' && s[i] == ')')
                {
                    stk.pop();
                    continue;
                }
                if (top == '[' && s[i] == ']')
                {
                    stk.pop();
                    continue;
                }
                if (top == '{' && s[i] == '}')
                {
                    stk.pop();
                    continue;
                }
            }
            return false;
        }
        return stk.size() == 0;
    }
};

//155
class MinStack
{
private:
    struct StackNode
    {
        int data;
        int min;
        StackNode *next;
    };
    StackNode *head;

public:
    /** initialize your data structure here. */
    MinStack()
    {
        this->head = new StackNode();
        this->head->next = NULL;
    }

    void push(int x)
    {

        StackNode *pushNode = new StackNode();
        pushNode->data = x;
        if (this->head->next == NULL)
        {
            pushNode->min = x;
        }
        else
        {
            pushNode->min = this->head->next->min > x ? x : this->head->next->min;
        }

        pushNode->next = this->head->next;
        this->head->next = pushNode;
    }

    void pop()
    {
        StackNode *delNode = this->head->next;
        this->head->next = delNode->next;
        delete delNode;
    }

    int top()
    {
        return this->head->next->data;
    }

    int getMin()
    {
        return this->head->next->min;
    }
};

int main()
{
    Solution s;
    bool isVa = s.isValid("");
    cout << "this string is " << isVa << endl;

    cout << "-------------------------------" << endl;

    MinStack minStack;
    minStack.push(-2);
    minStack.push(0);
    minStack.push(-1);
    int min = minStack.getMin();
    cout << "min data is " << min << endl;
    int top = minStack.top(); // return 0
    cout << "top number  " << top << endl;
    minStack.pop();
    int popMin = minStack.getMin();
    cout << "after pop min data is " << popMin << endl;
}
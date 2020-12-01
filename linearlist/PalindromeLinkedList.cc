/**
 * 234.Given a singly linked list, determine if it is a palindrome.
 * example 1 
 * Input: 1->2->2->1
 * Output: true
 * 
 * example 2
 * Input: 1->2
 * Output: false
 * O(n) time and O(1) space
 **/
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr){}
};

class Solution {
public:
    ListNode *left;
    bool isPalindrome(ListNode* head) {
        // ListNode *slow = head,*fast = head,*prev = nullptr;
        // //找到中心点  并且将前一段反转
        // while (fast !=nullptr && fast->next != nullptr)
        // {
        //     fast = fast->next->next;
        //     ListNode* next = slow->next;
        //     slow->next = prev;
        //     prev = slow;
        //     slow = next;  
        // }
        // if(fast != nullptr) slow = slow ->next;
        // head = prev; 
        // while (head && slow)
        // {
        //     if (head->val != slow->val) return false;
        //     head = head->next;
        //     slow = slow->next;
        // }
        // return true;


        //利用递归栈
        left = head;
        return traverse(head);
    }
    bool traverse(ListNode*right){
        if(right == nullptr)  return true;
        bool res = traverse(right->next);
        res = res && (right->val == left->val);
        left = left->next;
        return res;
    }
};
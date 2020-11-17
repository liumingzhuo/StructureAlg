/**
 * Definition for singly-linked list.
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 */
struct ListNode {
     int val;
     ListNode *next;
     ListNode(int x) : val(x), next(nullptr) {}
};
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        if(!head)return nullptr;
        ListNode *prev,*cur;
        prev = head;
        cur = head->next;
        while (cur)
        {
           ListNode * next = cur->next;
           cur->next = prev;
           prev = cur;
           cur = next;
        }
        head->next = nullptr;
        head = prev;
        return head;
    }
};
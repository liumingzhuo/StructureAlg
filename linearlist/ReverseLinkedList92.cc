/**
 * 92. Reverse Linked List II
 * 
 * Reverse a linked list from position m to n. Do it in one-pass.
 * 
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
 * */

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
};
class Solution {
public:
    ListNode* rearNode;
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        //base case
        if(m == 1)return reverseN(head, n);
        // 前进到反转的起点触发 base case
        head->next = reverseBetween(head->next,m - 1, n - 1);
        return head;
    }

    ListNode* reverseN(ListNode* head, int n){
        if(n == 1){
            rearNode = head->next;
            return head; 
        }
        ListNode* last = reverseN(head->next, n - 1);
        head->next->next = head;
        // 让反转之后的 head 节点和后面的节点连起来
        head->next = rearNode;
        return last;
    }
};
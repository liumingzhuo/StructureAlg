/**
 * 25. Reverse Nodes in k-Group
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
Example
给你这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->4->5
 * */

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};
class Solution {
public:
    ListNode* reverseKGroup(ListNode* head, int k) {
        if(head == nullptr) return nullptr;
        ListNode* start, *end;
        start = head, end = head;
        //移动end 到k位置
        for(int i = 0; i < k; i++){
            if(end == nullptr) return head;
            end = end->next;
        }
        //反转[start,end)
        ListNode *reverseNode = reverse(start,end);
        //拼接
        start->next = reverseKGroup(end, k);
        return reverseNode;
    }

    ListNode* reverse(ListNode *start, ListNode *end){
        ListNode* prev, *cur, *next;
        prev = nullptr, cur = start, next = start;
        while(cur != end){
            next = cur->next;
            cur ->next = prev;
            prev = cur;
            cur = next;
        } 
        return prev;
    }
};
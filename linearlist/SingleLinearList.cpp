#include <stdio.h>
#include <stdlib.h>
typedef struct LNode
{
    int data;
    struct LNode *next;
} LNode, *LinkList;

bool InitList(LinkList &l)
{
    l = (LNode *)malloc(sizeof(LNode));
    if (l == NULL)
        return false;
    l->next = NULL;
    return true;
}
//前插操作，在p结点前插入元素e
bool InsertPriorNode(LNode *p, int e)
{
    if (p == NULL)
        return false;
    LNode *s = (LNode *)malloc(sizeof(LNode));
    if (s == NULL)
        return false;
    s->next = p->next;
    p->next = s;
    s->data = p->data;
    p->data = e;
    return true;
}
//后插操作，在p结点后插入元素e
bool InsertNextNode(LNode *p, int e)
{
    if (p == NULL)
        return false;
    LNode *s = (LNode *)malloc(sizeof(LNode));
    if (s == NULL)
        return false;
    s->data = e;
    s->next = p->next;
    p->next = s;
    return true;
}

//按位插入 在i位插入e
bool InsertList(LinkList &l, int i, int e)
{
    if (i < 1)
        return false;
    int j = 0;
    LNode *p = l;
    while (p != NULL && j < i - 1)
    {
        p = p->next;
        j++;
    }
    InsertNextNode(p, e);
    return true;
}
//尾部初始化表
LinkList InitAfterList(LinkList &l)
{
    int x;
    l = (LinkList)malloc(sizeof(LNode));
    LNode *s, *r = l;
    scanf("%d", &x);
    while (x != 9999)
    {
        s = (LNode *)malloc(sizeof(LNode));
        s->data = x;
        r->next = s;
        r = s;
        scanf("%d", &x);
    }
    r->next = NULL;
    return l;
}

//头插法建立单链表
LinkList InitHeadersList(LinkList &l)
{
    int x;
    l = (LNode *)malloc(sizeof(LNode));
    l->next = NULL;
    LNode *s;
    scanf("%d", &x);
    while (x != 9999)
    {
        s = (LNode *)malloc(sizeof(LNode));
        s->data = x;
        s->next = l->next;
        l->next = s;
        scanf("%d", &x);
    }
    return l;
}
//删除指定节点
bool DeleteNode(LinkList &l, LNode *p)
{
    if (p == NULL)
        return false;
    LNode *q = p->next;

    if (q == NULL)
    {
        /*code*/
    }
    else
    {
        p->data = p->next->data;
        p->next = q->next;
        free(q);
    }
    return true;
}
//按位删除，并返回删除元素值
bool DeleteList(LinkList &l, int i, int &e)
{
    if (i < 1)
        return false;
    LNode *p = l;
    int j = 0;
    while (p != NULL && j < i - 1)
    {
        p = p->next;
        j++;
    }
    if (p == NULL)
        return false;
    LNode *q = p->next;
    if (q == NULL)
        return false;
    e = q->data;
    p->next = q->next;
    free(q);
    return true;
}
//按位查找
LNode *GetElem(LinkList &l, int i)
{
    if (i < 0)
        return NULL;
    LNode *p = l;
    int j = 0;
    while (p != NULL && j < i)
    {
        p = p->next;
        j++;
    }
    return p;
}
//按值查找
LNode *LocateElem(LinkList &l, int e)
{
    LNode *p = l->next;
    while (p != NULL && p->data != e)
    {
        p = p->next;
    }
    return p;
}
int LengthList(LinkList &l)
{
    int len = 0;
    while (l->next != NULL)
    {
        l->next = l->next->next;
        len++;
    }
    return len;
}
void PrintfList(LinkList &l)
{

    LNode *p = l->next;
    while (p != NULL)
    {
        printf("current element %d\n", p->data);
        p = p->next;
    }
}

//单链表反转
void ReverseList(LinkList &l)
{
    if (!l->next)
        return;
    LNode *pPrev, *pCur, *pNext;
    pPrev = l->next;
    if (!pPrev->next)
        return;
    pCur = pPrev->next;
    pPrev->next = NULL; //将第一个元素 与 表头断开
    while (pCur->next)
    {
        pNext = pCur->next;
        pCur->next = pPrev;
        pPrev = pCur;
        pCur = pNext;
    }
    pCur->next = pPrev; //将pCur->next 连接前面的链表
    l->next = pCur;
}
//两个有序的链表合并  leetcode 21
LNode *MergeList(LinkList &l1, LinkList &l2)
{
    if (l1 == NULL)
        return l2;
    if (l2 == NULL)
        return l1;
    if (l1->data <= l2->data)
    {
        l1->next = MergeList(l1->next, l2);
        return l1;
    }
    else
    {
        l2->next = MergeList(l1, l2->next);
        return l2;
    }
}
//删除链表倒数第n个元素 leetcode 19
LNode *removeNthFromEnd(LinkList &l, int n)
{
    if (l == NULL || l->next == NULL || n == 0)
        return NULL;
    //类似于快慢指针
    LNode *pPrev = l;
    LNode *pCur = l;
    while (n > 0)
    {
        pCur = pCur->next;
        --n;
    }
    if (pCur == NULL)
        return l->next;
    while (pCur->next)
    {
        pCur = pCur->next;
        pPrev = pPrev->next;
    }
    pPrev->next = pPrev->next->next;
    return l;
}

//删除链表倒数第n个元素 leetcode 19 递归
int cur = 0;
LNode *removeNthFromEnd2(LinkList &l, int n)
{
    if (!l)
        return NULL;
    l->next = removeNthFromEnd2(l->next, n);
    cur++;
    if (cur == n)
        return l->next;
    return l;
}

int main()
{
    LinkList l;
    // InitList(l);
    // l = InitAfterList(l);
    l = InitHeadersList(l);

    // ReverseList(l);
    removeNthFromEnd2(l, 2);
    PrintfList(l);
    // InsertList(l, 2, 2);

    // int deleteEle;
    // DeleteList(l, 1, deleteEle);
    // printf("删除的元素:%d\n", deleteEle);

    // LNode *findElem = GetElem(l, 2);
    // printf("查找第2位的值为:%d\n", findElem->data);

    // PrintfList(l);
    // int len = LengthList(l);
    // printf("length %d\n", len);
    return 0;
}

#include <stdio.h>
#include <stdlib.h>
typedef struct DNode
{
    int data;
    struct DNode *next, *prior;
} DNode, *DLinkedList;

bool InitDLinkedList(DLinkedList &l)
{
    DNode *p = (DNode *)malloc(sizeof(DNode));
    if (p == NULL)
        return false;
    p->next = NULL;
    p->prior = NULL;
    return true;
}

bool EmptyList(DLinkedList &l)
{
    return l->next == NULL;
}

//在p结点后插入结点s
bool InsertNextNode(DNode *p, DNode *s)
{
    if (p == NULL || s == NULL)
        return false;
    s->next = p->next;
    if (p->next != NULL)
    {
        p->next->prior = s;
    }
    p->next = s;
    s->prior = p;
    return true;
}
//删除p的后继结点
bool DeleteNextNode(DNode *p)
{
    if (p == NULL)
        return false;
    DNode *q = p->next;
    if (q == NULL)
        return false;
    p->next = q->next;
    if (q->next != NULL)
        q->next->prior = p;
    free(q);
    return true;
}
int main()
{
    DLinkedList l;
    InitDLinkedList(l);
    bool isEmpty = EmptyList(l);
    printf("链表是否为空%d\n", isEmpty);
}
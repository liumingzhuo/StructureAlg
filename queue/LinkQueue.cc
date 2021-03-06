#include <iostream>
using namespace std;

typedef int TypeEle;
typedef struct _queueNode
{
    TypeEle data;
    _queueNode *next;
} queueNode;
typedef struct _queueList
{
    int num;
    queueNode *head;
    queueNode *tail;
} queueList;
class LinkQueue
{
private:
    queueList *queue;

public:
    LinkQueue();
    ~LinkQueue();
    void enqueue(TypeEle e);
    TypeEle dequeue();
    int queueLength();
};

LinkQueue::LinkQueue()
{
    queue = (queueList *)malloc(sizeof(queueList));
    if (!queue)
        return;
    queue->num = 0;
    queue->tail = NULL;
    queue->head = NULL;
}

LinkQueue::~LinkQueue()
{
}
void LinkQueue::enqueue(TypeEle e)
{
    queueNode *newnode;
    newnode = (queueNode *)malloc(sizeof(queueNode));
    newnode->data = e;
    newnode->next = NULL;
    if (queue->head == NULL)
    {
        queue->head = newnode;
    }
    else
    {
        queue->tail->next = newnode;
    }
    queue->tail = newnode;
    queue->num++;
}

TypeEle LinkQueue::dequeue()
{
    if (queue == NULL || queue->num == 0)
    {
        return -99999;
    }
    queueNode *ptmp = NULL;
    TypeEle data = queue->head->data;
    ptmp = queue->head;
    queue->head = ptmp->next;
    if (queue->head == NULL)
    {
        queue->tail = NULL;
    }
    delete ptmp;
    return data;
}
int LinkQueue::queueLength()
{
    return queue->num;
}
int main()
{
    LinkQueue queue;
    queue.enqueue(1);
    queue.enqueue(2);
    queue.enqueue(3);
    cout << "queue length is " << queue.queueLength() << endl;
    TypeEle rlt = queue.dequeue();
    cout << "first dequeue " << rlt << endl;
    TypeEle rlt2 = queue.dequeue();
    cout << "second dequeue " << rlt2 << endl;
    queue.enqueue(4);
    cout << "last queue length is " << queue.queueLength() << endl;
}
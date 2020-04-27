/**
 * 
 * 循环队列
 * 
 **/
#include <iostream>
using namespace std;
typedef int TypeEle;

class CircularQueue
{
private:
    TypeEle *array;
    int head;
    int tail;
    int capacity;

public:
    CircularQueue(int n);
    ~CircularQueue();
    void enqueue(TypeEle e);
    TypeEle dequeue();
};

CircularQueue::CircularQueue(int n)
{
    array = (TypeEle *)malloc(sizeof(TypeEle) * n);
    head = 0;
    tail = 0;
    capacity = n;
}

CircularQueue::~CircularQueue()
{
    delete[] array;
}

void CircularQueue::enqueue(TypeEle e)
{
    if ((tail + 1) % capacity == head)
        return;
    array[tail] = e;
    tail = (tail + 1) % capacity;
}

TypeEle CircularQueue::dequeue()
{
    if (tail == head)
        return NULL;
    TypeEle rlt = array[head];
    head = (head + 1) % capacity;
    return rlt;
}

int main()
{
    CircularQueue queue = CircularQueue(8);
    queue.enqueue(1);
    queue.enqueue(2);
    queue.enqueue(3);
    int rlt1 = queue.dequeue();
    queue.enqueue(4);
    queue.enqueue(5);
    queue.enqueue(6);
    queue.enqueue(7);
    queue.enqueue(8);
    int rlt2 = queue.dequeue();
    cout << rlt1 << endl;
    cout << rlt2 << endl;
}

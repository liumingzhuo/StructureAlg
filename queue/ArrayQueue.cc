/**
 * 
 * 顺序队列
 * 
 **/

#include <iostream>
using namespace std;
template <class T>
class ArrayQueue
{
public:
    ArrayQueue(int n);
    ~ArrayQueue();
    void enqueue(T x);
    T dequeue();
    int queueLength();

private:
    T *array;
    int head;
    int tail;
    int capacity;
};

template <class T>
ArrayQueue<T>::ArrayQueue(int n)
{
    array = (T *)malloc(sizeof(T) * n);
    if (array == NULL)
        return;
    this->capacity = n;
    this->tail = 0;
    this->head = 0;
}
template <class T>
ArrayQueue<T>::~ArrayQueue()
{
    delete[] array;
}

template <class T>
void ArrayQueue<T>::enqueue(T x)
{
    if (this->tail == this->capacity)
    {

        if (this->head == 0)
        {
            T *newbase;
            newbase = (T *)realloc(array, sizeof(T) * this->capacity + this->capacity);
            this->array = newbase;
            this->capacity *= 2;
        }
        else
        {
            for (int i = this->head; i < this->tail; i++)
            {
                this->array[i - this->head] = this->array[i];
            }
            this->tail -= head;
            this->head = 0;
        }
    }
    this->array[this->tail] = x;
    this->tail++;
}

template <class T>
T ArrayQueue<T>::dequeue()
{
    if (this->head == this->tail)
        return -100000;
    T rlt = array[head];
    head++;
    return rlt;
}

template <class T>
int ArrayQueue<T>::queueLength()
{
    return tail;
}

int main()
{
    ArrayQueue<int> queue = ArrayQueue<int>(2);
    queue.enqueue(1);
    queue.enqueue(2);
    cout << queue.queueLength() << endl;
    queue.enqueue(3);
    cout << "扩容---->" << queue.queueLength() << endl;
    int rlt = queue.dequeue();
    cout << "出队列--->" << rlt << endl;
}

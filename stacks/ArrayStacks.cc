
#include <iostream>
using namespace std;
template <class T>
class ArrayStack
{
public:
    ArrayStack(int n);
    ~ArrayStack();
    void push(T e);
    T pop();
    void printStack();

private:
    T *data;
    int count;
    int n;
};
template <class T>
ArrayStack<T>::ArrayStack(int stackN)
{
    n = stackN;
    data = new T[n];
    count = 0;
}
template <class T>
ArrayStack<T>::~ArrayStack()
{
    delete[] data;
}

template <class T>
void ArrayStack<T>::push(T e)
{
    if (count == n)
        return;
    data[count++] = e;
}
template <class T>
T ArrayStack<T>::pop()
{
    return data[count--] ? count != -1 : NULL;
}
template <class T>
void ArrayStack<T>::printStack()
{
    for (int i = 0; i < count; i++)
    {
        cout << data[i] << endl;
    }
}

int main()
{
    ArrayStack<int> as(10);
    as.push(1);
    as.push(2);
    as.push(3);
    as.push(4);
    as.push(5);
    as.printStack();
    cout << "-------------" << endl;
    as.pop();
    as.pop();
    as.printStack();
}
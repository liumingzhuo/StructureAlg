#include <iostream>

using namespace std;
template <class T>
struct Node
{
    T data;
    Node<T> *next;
    Node(T e)
    {
        this->data = e;
        next = NULL;
    };
    Node()
    {
        next = NULL;
    }
};

template <class T>
class LinkListStack
{
public:
    LinkListStack();
    ~LinkListStack();
    int size();
    void push(T e);
    T pop();
    T top();

private:
    Node<T> *head;
    int count;
};
template <class T>
LinkListStack<T>::LinkListStack()
{
    this->head = new Node<T>();
    this->head->next = NULL;
    this->count = 0;
}
template <class T>
LinkListStack<T>::~LinkListStack()
{
    Node<T> *ptr, *temp;
    ptr = head;
    while (ptr->next != NULL)
    {
        temp = ptr->next;
        ptr->next = temp->next->next;
        delete temp;
        count--;
    }
    delete head;
    this->head = NULL;
    this->count = 0;
}
template <class T>
void LinkListStack<T>::push(T e)
{
    Node<T> *insertNode = new Node<T>(e);
    insertNode->next = head->next;
    head->next = insertNode;
    count++;
    cout << "push data: " << head->next->data << endl;
}
template <class T>
T LinkListStack<T>::pop()
{
    if (head->next == NULL || count == 0)
        return -1;
    Node<T> *temp = head->next;
    head->next = temp->next;
    T data = temp->data;
    delete temp;
    count--;
    return data;
}
template <class T>
T LinkListStack<T>::top()
{
    if (head->next == NULL || count == 0)
        return -1;
    return head->next->data;
}
template <class T>
int LinkListStack<T>::size()
{
    return count;
}

int main()
{
    LinkListStack<int> ls;
    ls.push(1);
    ls.push(2);
    int ss = ls.size();
    cout << "current size is " << ss << endl;
    int delData = ls.pop();
    cout << "pop number is " << delData << endl;
    int topData = ls.top();
    cout << "top number is " << topData << endl;
}

#include <iostream>
using namespace std;

template <class T>
int length(T &arr)
{
    return sizeof(arr) / sizeof(arr[0]);
}

int SumArray(int arr[], int len)
{
    if (len == 0)
        return 0;
    if (len == 1)
        return arr[len - 1];
    return arr[len - 1] + SumArray(arr, len - 1);
}

int main()
{
    int array[] = {2, 4, 6};
    int rlt = SumArray(array, length(array));
    cout << rlt << endl;
}
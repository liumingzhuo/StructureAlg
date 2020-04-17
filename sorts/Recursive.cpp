#include <iostream>
#include <stdio.h>
using namespace std;

template <class T>
int length(T &arr)
{
    return sizeof(arr) / sizeof(arr[0]);
}
//计算列表中元素的总和
int SumArray(int arr[], int len)
{
    if (len == 0)
        return 0;
    if (len == 1)
        return arr[len - 1];
    return arr[len - 1] + SumArray(arr, len - 1);
}

//找出数组中最大的数字
int MaxNumber(int arr[], int len)
{
    int temp;
    if (len == 0)
        return 0;
    if (len == 1)
        return arr[0];
    temp = MaxNumber(arr, len - 1);
    return temp > MaxNumber(arr, len - 1) ? temp : arr[len - 1]; //返回值的元素和当前元素做对比
}

int main()
{
    int array[] = {3, 2, 6};
    int len = length(array);
    int rlt = SumArray(array, len);
    cout << "sum array number is " << rlt << endl;
    int MaxNum = MaxNumber(array, len);
    cout << "max number is " << MaxNum << endl;
}
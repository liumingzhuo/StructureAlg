/**
 * 
 * 递归思想
 * 
 * */
#include <iostream>
#include <stdio.h>
using namespace std;

template <class T>
int length(T &arr)
{
    return sizeof(arr) / sizeof(arr[0]);
}
//1.计算列表中元素的总和
int SumArray(int arr[], int len)
{
    if (len == 0)
        return 0;
    if (len == 1)
        return arr[len - 1];
    return arr[len - 1] + SumArray(arr, len - 1);
}

//2.找出数组中最大的数字
int MaxNumber(int *arr, int index)
{
    if (index == 0)
        return arr[0];
    return arr[index] > MaxNumber(arr, index - 1) ? arr[index] : MaxNumber(arr, index - 1);
}

int main()
{
    int array[] = {3, 2, 6};
    int len = length(array);
    int rlt = SumArray(array, len);
    cout << "sum array number is " << rlt << endl;
    int MaxNum = MaxNumber(array, len - 1);
    cout << "max number is " << MaxNum << endl;
}
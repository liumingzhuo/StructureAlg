
#include <iostream>
using namespace std;
/**
 * 
 * select sort
 * O(n^2)
 * 
 * */
template <class T>
int length(T &arr)
{
    return sizeof(arr) / sizeof(arr[0]);
}

void SelectSort(int *arr, int len)
{
    int index = 0;
    for (int i = 0; i < len - 1; i++)
    {
        //将最小的数插入到数组第一个位置，j可以从i+1个位置开始遍历
        for (int j = i + 1; j < len; j++)
        {
            if (arr[j] < arr[index])
                index = j;
        }
        if (arr[index] < arr[i])
        {
            int temp = arr[i];
            arr[i] = arr[index];
            arr[index] = temp;
        }
    }
}

int main()
{
    int arr[] = {5, 3, 6, 2, 10};
    int len = length(arr);
    SelectSort(arr, len);
    for (int i = 0; i < len; i++)
    {
        cout << arr[i] << endl;
    }
}

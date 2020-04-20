/**
 * 
 * quick sort
 * divide and conquer(D&C)
 * 
 * */
#include <iostream>
#include <vector>
using namespace std;

int partition(vector<int> &vi, int low, int high)
{
    int flag = low - 1;
    int pivot = vi[high];
    for (int j = low; j < high; j++)
    {
        if (vi[j] < pivot)
        {
            flag++;
            swap(vi[flag], vi[j]);
        }
    }
    swap(vi[flag + 1], vi[high]);
    return flag + 1;
}

void quickSort(vector<int> &vi, int low, int high)
{
    if (low < high)
    {
        int mid = partition(vi, low, high);
        quickSort(vi, low, mid - 1);
        quickSort(vi, mid + 1, high);
    }
}

int main()
{
    int arr[] = {10, 5, 2, 3, 2, 10};
    vector<int> vi(arr, arr + 6);
    quickSort(vi, 0, vi.size() - 1);
    for (int v : vi)
        cout << v << " ";
}
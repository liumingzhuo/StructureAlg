//31. 下一个排列
package main

//从后向前找 找到a[i] < a[i+1]
//再找个 a[j] > a[i]
//swap(a[j],a[i])
//此时a[i+1...n-1] 是降序， 重排成升序
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(&nums, i+1)
}
func reverse(nums *[]int, start int) {
	left, right := start, len(*nums)-1
	for left < right {
		(*nums)[left], (*nums)[right] = (*nums)[right], (*nums)[left]
		left++
		right--
	}
}
func main() {
	nextPermutation([]int{1, 2, 3})
}

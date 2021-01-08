/*
283. 移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
example:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
*/
package MoveZeroes283

func moveZeroes(nums []int) {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
}

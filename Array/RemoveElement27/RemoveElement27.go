/*
27. 移除元素
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/
package RemoveElement27

func removeElement(nums []int, val int) int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		//当fast不同于val时  就让slow等于fast的值
		if nums[fast] != val {
			//需要先赋值给nums[slow]   这样可以保证nums[0....slow-1]是不包含val值的
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	//因为前面nums移除后的长度为slow-1 那么slow就是新数组的长度
	return slow
}

package finderrornums645

//645错误的集合

/*
输入：nums = [1,2,2,4]
输出：[2,3]
*/
func findErrorNums(nums []int) []int {
	n := len(nums)
	tmp := make([]int, n+1) //初始化tmp切片
	rlt1, rlt2 := 0, 0
	for i := 0; i < n; i++ {
		tmp[nums[i]]++
	}
	for i := 1; i <= n; i++ {
		if tmp[i] == 0 { //为0 则是缺失
			rlt1 = i
		} else if tmp[i] == 2 { //为2则表示出现了2次
			rlt2 = i
		}
	}
	return []int{rlt2, rlt1}
}

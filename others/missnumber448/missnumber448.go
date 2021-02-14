package missnumber448

//448. 找到所有数组中消失的数字   [1,n]
func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			res = append(res, i+1)
		}
	}
	return res
}

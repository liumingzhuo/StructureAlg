package subsets78

/*78. 子集*/
var res [][]int

func subsets(nums []int) [][]int {
	res = [][]int{}
	path := []int{}
	backtrack(nums, 0, path)
	return res
}
func backtrack(nums []int, start int, path []int) {
	//因为go语言切片共用底层数组，所以需要copy
	temp := make([]int, len(path))
	copy(temp, path)
	res = append(res, temp)
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrack(nums, i+1, path)
		path = path[:len(path)-1]
	}
}

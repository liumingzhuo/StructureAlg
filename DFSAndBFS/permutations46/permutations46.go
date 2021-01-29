package permutations46

/*
46. 全排列
*/

var res [][]int

func permute(nums []int) [][]int {
	// for i := 0; i < len(res); i++ {
	// 	res[i] = make([]int, 0)
	// }
	res = [][]int{}
	track := []int{}
	backtrack(nums, track)
	return res
}
func backtrack(nums []int, track []int) {
	//回朔结束
	if len(track) == len(nums) {
		//切片底层公用数组，需要copy
		tmp := make([]int, len(nums))
		copy(tmp, track)
		res = append(res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		//排除重复
		exit := false
		for _, v := range track {
			if v == nums[i] {
				exit = true
				break
			}
		}
		if exit {
			continue
		}
		track = append(track, nums[i])
		backtrack(nums, track)
		track = track[:len(track)-1]
	}
}

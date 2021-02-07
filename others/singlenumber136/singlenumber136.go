package singlenumber136

//[2,2,1]
//1

//[4,1,2,1,2]
//4

//一个数和它本身做异或运算结果为 0，而和0做异或还是本身
func singleNumber(nums []int) int {
	rlt := 0
	for _, num := range nums {
		rlt ^= num
	}
	return rlt
}

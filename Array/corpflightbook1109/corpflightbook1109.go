package corpflightbook1109

//1109. 航班预订统计
//差分数组：频繁对原始数组的某个区间的元素进行增减
func corpFlightBookings(bookings [][]int, n int) []int {
	//差分数组
	diff := make([]int, n)
	for _, b := range bookings {
		i := b[0] - 1 //参数1
		j := b[1] - 1 //参数2
		val := b[2]
		//对区间增减val
		diff[i] += val
		if j+1 < len(diff) {
			diff[j+1] -= val
		}
	}
	res := make([]int, len(diff))
	res[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		res[i] = res[i-1] + diff[i]
	}
	return res

}

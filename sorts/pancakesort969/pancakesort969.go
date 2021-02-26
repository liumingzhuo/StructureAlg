package pancakesort969

//969 煎饼排序
var res []int

func pancakeSort(arr []int) []int {
	res = []int{}
	sort(arr, len(arr))
	return res
}

func sort(arr []int, n int) {
	if n == 1 {
		return
	}
	//最大饼
	maxCake := 0
	maxCakeIndex := 0
	for i := 0; i < n; i++ {
		if arr[i] > maxCake {
			maxCakeIndex = i
			maxCake = arr[i]
		}
	}
	reverseCake(arr, 0, maxCakeIndex)
	res = append(res, maxCakeIndex+1)
	reverseCake(arr, 0, n-1)
	res = append(res, n)
	sort(arr, n-1)
}
func reverseCake(arr []int, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

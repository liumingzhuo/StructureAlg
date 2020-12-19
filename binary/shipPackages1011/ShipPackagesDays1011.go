/*
1011. 在 D 天内送达包裹的能力
传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。

传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。

返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。
*/
package shipPackages

func shipWithinDays(weights []int, D int) int {
	left := getMax(weights)
	right := getSum(weights) + 1
	for left < right {
		mid := left + (right-left)/2
		if canFinish(weights, D, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func getMax(weights []int) int {
	var max int = 0
	for _, val := range weights {
		if val > max {
			max = val
		}
	}
	return max
}
func getSum(weights []int) int {
	var sum int = 0
	for _, val := range weights {
		sum += val
	}
	return sum
}

//D天内能完成的情况
func canFinish(weights []int, D int, cap int) bool {
	i := 0
	for day := 0; day < D; day++ {
		maxCap := cap
		for maxCap-weights[i] >= 0 {
			maxCap -= weights[i]
			i++
			if i == len(weights) {
				return true
			}
		}
	}
	return false
}

package poweroftwo231

/*
Input: n = 1
Output: true
Explanation: 20 = 1

Input: n = 16
Output: true
Explanation: 24 = 16

Input: n = 3
Output: false

Input: n = 4
Output: true

Input: n = 5
Output: false
*/

//一个数 如果是2的指数，那么二进制一定只有一个1
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

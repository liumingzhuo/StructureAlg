package countprimes204

/*
输入：n = 10
输出：4
解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。

输入：n = 0
输出：0

输入：n = 1
输出：0
*/

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}

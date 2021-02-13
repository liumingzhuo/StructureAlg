package superpow372

//372. 超级次方
const base int = 1337

// a^[1,2,3] = a^3 * (a^[1,2])^10
func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]
	part1 := myPow(a, last)
	part2 := myPow(superPow(a, b), 10)
	return (part1 * part2) % base
}

//(a*b)%c   =  (a%c)(b%c)%c
// a^b = a* a^(b-1)// b为奇数
// a^b = (a^(b/2))^2 b为偶数
func myPow(a int, k int) int {
	if k == 0 {
		return 1
	}
	a = a % base
	if k%2 == 1 {
		return (a * myPow(a, k-1)) % base
	} else {
		sub := myPow(a, k/2)
		return (sub * sub) % base
	}
}

//(a*b)%c   =  (a%c)(b%c)%c
/*
func myPow(a int, k int) int {
	a = a % base
	res := 1
	for i := 0; i < k; i++ {
		res *= a
		res %= base
	}
	return res
}
*/

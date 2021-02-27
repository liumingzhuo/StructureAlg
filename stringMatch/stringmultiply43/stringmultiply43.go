package stringmultiply43

//43. 字符串相乘
//输入: num1 = "2", num2 = "3"
//输出: "6"
func multiply(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	res := make([]int, n1+n2)
	num1b := []byte(num1)
	num2b := []byte(num2)
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			mul := (num1b[i] - '0') * (num2b[j] - '0')
			p1 := i + j
			p2 := i + j + 1
			total := int(mul) + res[p2]
			res[p2] = total % 10
			res[p1] += total / 10
		}
	}
	var i int
	for ; i < len(res) && res[i] == 0; i++ {
	}
	var bs []byte
	for ; i < len(res); i++ {
		bs = append(bs, byte(res[i])+'0')
	}
	if len(bs) > 0 {
		return string(bs)
	} else {
		return "0"
	}

}

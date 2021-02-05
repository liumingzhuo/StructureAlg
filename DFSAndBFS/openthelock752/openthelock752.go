package openthelock752

//752. 打开转盘锁
//Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
//Output: 6

//Input: deadends = ["8888"], target = "0009"
//Output: 1

//Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
//Output: -1

//往上拨
func up(s string, j int) string {
	b := []byte(s)
	if b[j] == '9' {
		b[j] = '0'
	} else {
		b[j]++
	}
	return string(b)
}

//往下拨
func down(s string, j int) string {
	b := []byte(s)
	if b[j] == '0' {
		b[j] = '9'
	} else {
		b[j]--
	}
	return string(b)

}
func openLock(deadends []string, target string) int {
	visited := make(map[string]int) //记录算过的 以免死循环
	deads := make(map[string]int, len(deadends))
	for _, v := range deadends {
		deads[v] = 1
	}

	queue := make([]string, 0)
	queue = append(queue, "0000")
	visited["0000"] = 1
	step := 0
	for len(queue) != 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			//如果出现死亡数字 那么跳过 不能继续拨了
			if _, ok := deads[cur]; ok {
				continue
			}
			//遇到目标数字 返回
			if cur == target {
				return step
			}
			//拨动 4个位置的每个字符
			for j := 0; j < 4; j++ {
				strUp := up(cur, j)
				if _, ok := visited[strUp]; !ok {
					queue = append(queue, strUp)
					visited[strUp] = 1
				}
				strDown := down(cur, j)
				if _, ok := visited[strDown]; !ok {
					queue = append(queue, strDown)
					visited[strDown] = 1
				}
			}
		}
		step++
	}
	return -1
}

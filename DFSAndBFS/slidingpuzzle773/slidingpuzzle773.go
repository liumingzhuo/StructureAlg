package slidingpuzzle773

//输入：board = [[1,2,3],[4,0,5]]
//输出：1   解释：交换 0 和 5 ，1 步完成

//输入：board = [[1,2,3],[5,4,0]]
//输出：-1 解释：没有办法完成谜板

//输入：board = [[4,1,2],[5,0,3]]
//输出：5

//输入：board = [[3,2,4],[1,5,0]]
//输出：14

//BFS
func slidingPuzzle(board [][]int) int {
	//将二维数组 转成字符串
	baseStr := ""
	m, n := 2, 3
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			baseStr += string(byte(board[i][j] + '0'))
		}
	}
	//2*3的数组  下标对应的上下左右的索引,是固定的, 即0的下标对应右边是1  下边是3
	neighbor := [][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}
	target := "123450"
	step := 0
	visited := make(map[string]struct{}, 0)
	visited[baseStr] = struct{}{}
	q := []string{}
	q = append(q, baseStr)
	for len(q) != 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if cur == target {
				return step
			}
			//找到0对应的下标
			idx := 0
			for cur[idx] != '0' {
				idx++
			}
			for _, v := range neighbor[idx] {
				//扩散 下标v
				bbb := []byte(cur)
				bbb[idx], bbb[v] = bbb[v], bbb[idx]
				newStr := string(bbb)
				if _, ok := visited[newStr]; !ok {
					q = append(q, newStr)
					visited[newStr] = struct{}{}
				}
			}
		}
		step++
	}
	return -1
}

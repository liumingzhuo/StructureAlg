/*
岛屿数量
DFS
*/

package islands200

func numIslands(grid [][]byte) int {
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				helper(grid, i, j)
				ret++
			}
		}
	}
	return ret
}

func helper(grid [][]byte, x, y int) {
	if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[x]) && grid[x][y] == '1' {
		grid[x][y] = '0'
		helper(grid, x, y-1)
		helper(grid, x, y+1)
		helper(grid, x-1, y)
		helper(grid, x+1, y)
	}
}

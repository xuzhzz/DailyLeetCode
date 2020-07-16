package dailyleetcode

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start

var h, w int

// Pos index pos
type Pos struct {
	x int
	y int
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	lands := 0
	h = len(grid)
	w = len(grid[0])
	for i := 0; i < h; i = i + 1 {
		for j := 0; j < w; j = j + 1 {
			if grid[i][j] == '1' {
				visitIsland(grid, i, j)
				lands = lands + 1
			}
		}
	}
	return lands
}

func visitIsland(grid [][]byte, x, y int) {
	queue := make([]Pos, 0)
	queue = append(queue, Pos{x, y})
	for len(queue) > 0 {
		x = queue[0].x
		y = queue[0].y
		if y-1 >= 0 && grid[x][y-1] == '1' {
			grid[x][y-1] = '0'
			queue = append(queue, Pos{x, y - 1})
		}
		if y+1 < w && grid[x][y+1] == '1' {
			grid[x][y+1] = '0'
			queue = append(queue, Pos{x, y + 1})
		}
		if x-1 >= 0 && grid[x-1][y] == '1' {
			grid[x-1][y] = '0'
			queue = append(queue, Pos{x - 1, y})
		}
		if x+1 < h && grid[x+1][y] == '1' {
			grid[x+1][y] = '0'
			queue = append(queue, Pos{x + 1, y})
		}
		queue = queue[1:]
	}
	return
}

// @lc code=end

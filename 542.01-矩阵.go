package dailyleetcode

/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
type pos struct {
	x, y int
}

func updateMatrix(matrix [][]int) [][]int {
	res := make([][]int, 0)
	queue := make([][]int, 0)
	h := len(matrix)
	w := len(matrix[0])
	for i := 0; i < h; i++ {
		res = append(res, matrix[i])
		for j := 0; j < w; j++ {
			if matrix[i][j] == 0 {
				res[i][j] = 0
				queue = append(queue, []int{i, j})
			} else {
				res[i][j] = -1
			}
		}
	}
	offset := [][]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}

	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		for _, ost := range offset {
			visitX := tmp[0] + ost[0]
			visitY := tmp[1] + ost[1]
			if visitX < 0 || visitY < 0 || visitX >= h || visitY >= w || res[visitX][visitY] == 0 {
				continue
			}
			if res[visitX][visitY] == -1 {
				res[visitX][visitY] = res[tmp[0]][tmp[1]] + 1
				queue = append(queue, []int{visitX, visitY})
			}
		}
	}
	// BFS
	return res
}

// @lc code=end

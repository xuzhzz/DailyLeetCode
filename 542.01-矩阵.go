package dailyleetcode

/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(matrix [][]int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		res = append(res, matrix[i])
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				res[i][j] = 0
			} else {
				res[i][j] = -1
			}
		}
	}

	// BFS
	return res
}

// @lc code=end

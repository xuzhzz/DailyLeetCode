package dailyleetcode

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	largestArea := 0
	for i := 0; i < len(heights); i++ {
		area := findWidth(heights, i) * heights[i]
		if area > largestArea {
			largestArea = area
		}
	}
	return largestArea
}

func findWidth(heights []int, idx int) (width int) {
	width = 1
	for i := idx - 1; i >= 0 && heights[i] >= heights[idx]; i-- {
		width++
	}
	for i := idx + 1; i < len(heights) && heights[i] >= heights[idx]; i++ {
		width++
	}
	return width
}

// @lc code=end

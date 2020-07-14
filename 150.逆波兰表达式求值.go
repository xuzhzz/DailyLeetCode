package dailyleetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func isDigit(token string) bool {
	switch token {
	case "+", "-", "*", "/":
		return false
	default:
		return true
	}
}

func tranInt(token string) (res int) {
	res, _ = strconv.Atoi(token)
	return
}

func evalRes(num1, num2 int, symbol string) int {
	switch symbol {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 != 0 {
			return num1 / num2
		}
	}
	return 0
}

func evalRPN(tokens []string) int {
	if len(tokens) == 1 && isDigit(tokens[0]) {
		return tranInt(tokens[0])
	}
	if len(tokens) < 3 {
		return 0
	}
	stack := make([]int, 0)
	var result int
	for _, token := range tokens {
		if isDigit(token) {
			stack = append(stack, tranInt(token))
		} else {
			result = evalRes(stack[len(stack)-2], stack[len(stack)-1], token)
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		}
	}
	return result
}

// @lc code=end

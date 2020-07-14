package dailyleetcode

import "math"

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	Val    []int
	MinVal int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Val:    make([]int, 0),
		MinVal: int(math.MaxInt64),
	}
}

func (this *MinStack) Push(x int) {
	if this != nil {
		this.Val = append(this.Val, x)
		if x < this.MinVal {
			this.MinVal = x
		}
	}
}

func (this *MinStack) Pop() {
	if this != nil {
		if this.MinVal == this.Val[len(this.Val)-1] {
			this.MinVal = int(math.MaxInt64)
			for _, val := range this.Val[:len(this.Val)-1] {
				if val < this.MinVal {
					this.MinVal = val
				}
			}
		}
		this.Val = this.Val[:len(this.Val)-1]
	}
}

func (this *MinStack) Top() int {
	if this != nil {
		return this.Val[len(this.Val)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this != nil {
		return this.MinVal
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

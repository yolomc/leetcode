/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	Val  int
	Min  int
	Next *MinStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{0, 0, nil}
}

func (this *MinStack) Push(x int) {
	temp := *this
	min := x

	if this.Next != nil && min > this.Min {
		min = this.Min
	}

	this.Val = x
	this.Min = min
	this.Next = &temp
}

func (this *MinStack) Pop() {
	if this.Next != nil {
		this.Val = this.Next.Val
		this.Min = this.Next.Min
		this.Next = this.Next.Next
	} else {
		this.Val = 0
		this.Min = 0
		this.Next = nil
	}
}

func (this *MinStack) Top() int {
	return this.Val
}

func (this *MinStack) GetMin() int {
	return this.Min
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

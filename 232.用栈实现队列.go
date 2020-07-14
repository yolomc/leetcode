/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	s1 []int
	s2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var ret int
	if len(this.s2) > 0 {
		ret = this.s2[len(this.s2)-1]
		this.s2 = this.s2[:len(this.s2)-1]
	} else {
		for len(this.s1) > 1 {
			this.s2 = append(this.s2, this.s1[len(this.s1)-1])
			this.s1 = this.s1[:len(this.s1)-1]
		}
		ret = this.s1[0]
		this.s1 = this.s1[0:0]
	}
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {

	if this.Empty() {
		return 0
	}

	if len(this.s2) > 0 {
		return this.s2[len(this.s2)-1]
	} else {
		return this.s1[0]
	}
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

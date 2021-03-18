package Leet_Code

// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 栈 先进后出
// 一个队列中包含两个栈"先进后出的"栈 stack1&stack2
// 使用2个栈，来实现一个"先出先进的"队列

// 队列入列时，OutputStack 为空，InputStack满
// 队列出列时，InputStack 为空, OutputStack满


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

type CQueue struct {
	InputStack []int
	OutputStack []int
}


func Constructor() CQueue {
	queue := CQueue{[]int{}, []int{}}
	return queue
}


func (this *CQueue) AppendTail(value int)  {
	this.InputStack = append(this.InputStack , value)
}


func (this *CQueue) DeleteHead() int {

	if len(this.OutputStack)!= 0{
		x := this.OutputStack[len(this.OutputStack)-1]
		this.OutputStack = this.OutputStack[:len(this.OutputStack)-1]
		return x
	}
	if len(this.InputStack) != 0 {
		//  将InputStack中的依次pop到OutputStack中
		for i:= len(this.InputStack)-1 ;i>=0; i-- {
			this.OutputStack = append(this.OutputStack, this.InputStack[i])
			this.InputStack = this.InputStack[:len(this.InputStack)-1]
		}
		// 在从OutputStack
		x := this.OutputStack[len(this.OutputStack)-1]
		this.OutputStack = this.OutputStack[:len(this.OutputStack)-1]
		return x
	}
		// 都为-1
	return -1

}

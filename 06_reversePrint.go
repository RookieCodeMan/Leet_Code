package Leet_Code

// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
// 示例 1：
//
// 输入：head = [1,3,2]
// 输出：[2,3,1]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

//数组是一个固定长度的特定元素组成的序列，长度固定
//切片是可以增长和收缩动态序列，功能灵活
// 1.定义一个切片，让切片去引用已经创建好的数组
// 2.通过make方式创建切片
// 3.定义一个切片，直接指定具体的数据
func reversePrint(head *ListNode) []int {

	var valList  = [] int{} //定义空切片
	var reverseValList = [] int{}

	for {
		if head == nil{
			break
		} else {
			valList = append(valList, head.Val)
			head = head.Next
		}
	}
	valListLen := len(valList)
	// 法一，遍历链表
	for i:=0 ; i< valListLen; i++{
		reverseValList = append(reverseValList, valList[valListLen-i-1])
	}
	return reverseValList
}
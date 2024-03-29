package Leet_Code

// BM16 删除有序链表中重复的元素-I
//描述
//删除给出链表中的重复元素（链表中元素从小到大有序），使链表中的所有元素都只出现一次
//例如：
//给出的链表为1→1→2,返回1→2.
//给出的链表为1→1→2→3→3,返回1→2→3.
//
//数据范围：链表长度满足 1000≤n≤100，链表中任意节点的值满足 ∣val∣≤100
//进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

//解题思路: 核心代码 cur.Next =  cur.Next.Next
func deleteDuplicates_1( head *ListNode ) *ListNode {
	cur := head
	if cur == nil {
		return nil
	}
	for cur != nil {
		if cur.Next != nil {
			if cur.Val == cur.Next.Val {
				cur.Next = cur.Next.Next
				continue
			}
		}
		// 循环条件
		cur = cur.Next
	}
	return head
}

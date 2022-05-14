package Leet_Code

// BM16 删除有序链表中重复的元素-II
//描述
//给出一个升序排序的链表，删除链表中的所有重复出现的元素，只保留原链表中只出现一次的元素。
//例如：
//给出的链表为1→2→3→3→4→4→5, 返回1→2→5.
//给出的链表为1→1→1→2→3, 返回2→3.
//
//数据范围：链表长度 100000≤n≤10000，链表中的值满足 ∣val∣≤1000
//要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
//进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

// 解题思路：
// 双指针解法，设置一个newHead对象，pre = newHead，cur = head
// 然后循环判断cur对象，前后两个指针相差一个元素
func deleteDuplicates( head *ListNode ) *ListNode {
	newHead := &ListNode{0, nil}
	newHead.Next = head
	pre := newHead
	cur := head
	count := 0
	if cur == nil {
		return nil
	}
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			count ++
			// 有重复元素出现，就一直往后统计
			cur.Next = cur.Next.Next
		} else {
			if count > 0 {
				pre.Next = cur.Next
				count = 0 // 将count置为0
			} else {
				pre = cur
			}
			// 递归条件
			cur = cur.Next
		}
	}
	//存在比如{1，1，1}，因为删除，所以上述循环条件不进行判断，在此额外进行判断
	if count > 0 {
		pre.Next = cur.Next
	}

	return newHead.Next
}

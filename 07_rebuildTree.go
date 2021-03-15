package Leet_Code


// 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
// 例如，给出
//
// 前序遍历 preorder = [3,9,20,15,7]  中-左-右
// 中序遍历 inorder = [9,3,15,20,7]   左-中-右
// 返回如下的二叉树：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//  
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 思路，递归, 递归结束要有判断，否则将一直循环下去
//
// rootValue = 3
// 前序遍历 preorder = [3,9,20,15,7]  中-左-右
// 中序遍历 inorder =  [9,3,15,20,7]   左-中-右
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootValue := preorder[0]
	var rootIndex int

	//
	for i:=0;i<len(inorder);i++{
		if rootValue == inorder[i] {
			rootIndex = i
		}
	}

	left := buildTree(preorder[1:rootIndex+1],inorder[0:rootIndex])
	right := buildTree(preorder[rootIndex+1:],inorder[rootIndex+1:])
	return &TreeNode{rootValue,left, right}

}


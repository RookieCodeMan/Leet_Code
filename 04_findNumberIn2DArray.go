package Leet_Code

// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//
//
// 示例:
//
// 现有矩阵 matrix 如下：
//
// [
// [1,   4,  7, 11, 15],
// [2,   5,  8, 12, 19],
// [3,   6,  9, 16, 22],
// [10, 13, 14, 17, 24],
// [18, 21, 23, 26, 30]
// ]
// 给定 target = 5，返回 true。
//
// 给定 target = 20，返回 false

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	// 矩阵其实就像是一个Binary Search Tree，二分法查找(前提是该数组已经排序过了)
	if len(matrix) == 0 {
		return false
	}
	rowNum := len(matrix)
	colNum := len(matrix[0])
	r := 0
	c := colNum-1
	for ;r < rowNum && c>= 0;{
		num := matrix[r][c]
		if num == target{
			return true
		} else if num > target{
			c--
		} else {
			r++
		}
	}
	return false

}
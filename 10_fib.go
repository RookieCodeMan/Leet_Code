package Leet_Code
// 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
//
// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
// 斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//
// 示例 1：
// 输入：n = 2
// 输出：1

// 示例 2：
// 输入：n = 5
// 输出：5



// 方法1，直接递归，太多重复计算，效率太低
func Fib(n int) int {
	if n ==0 {
		return 0
	}
	if n ==1{
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// 方法2，重复计算太多，尽量避免重复计算，从下往上算，时间复杂度为O（n）
func Fib2(n int) int {
	if n ==0 || n ==1 {
		return n
	} else {
		fibPre := 0
		fibCur := 1
		var fibNext int
		for i:=2;i<=n;i++{
			fibNext = (fibPre+fibCur) % 1000000007
			fibPre = fibCur
			fibCur = fibNext
		}
		return fibNext
	}
}

// 方法3，动态规划
func Fib3(n int) int {
	if n ==0 || n ==1 {
		return n
	} else {
		dp := make([]int,n+1)
		dp[0] = 0
		dp[1] = 1
		for i:=2;i<=n;i++{
			dp[i] = (dp[i-1] + dp[i-2])% 1000000007
		}
		return dp[n]
	}
}
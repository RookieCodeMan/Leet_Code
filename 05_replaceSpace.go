package Leet_Code

import (
	"strings"
)

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"
//
// 示例 1：
//
// 输入：s = "We are happy."
// 输出："We%20are%20happy."

func ReplaceSpace(s string) string {
	// go常用处理函数
	//1.查找一个字符串再另一个字符串是否出现 strings.Contains
	//2.字符串的连接 strings.Join
	//3.查找一个字符串在另一个字符串第一次出现的位置下标，没有找到返回-1
	//4.Repeat 将一个字符串重复m次 Replace 字符串替换
	//5.Split 将一个字符串按照标志位分割
	//6.Trim 去掉字符串头尾的内容
	//7.Fields 去掉空格返回切片
	str := "%20"
	newString := strings.Replace(s,  " ", str, -1 )
	return newString
}

package main

import "fmt"

// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

// 示例 1：
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。

// 提示：
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成

func main() {
	strs := []string{"long", "longboard", "longevity"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	//[]string 数组为空直接范围空字符串 ""
	if len(strs) == 0 {
		return ""
	}

	//拿出数组第一个然后拿出第一个字母，依次开始比较
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			//i == len(strs[j]) len 遇到最短，如果不加入这个则会 i列 > j行 时候 panic
			//然后开始判断每一列比较同列上一位是否相同，相同则继续往下比较，不同则返回列之前部分
			//因为纵向比较，所以返回之前部分直接用第一行，不同之前的
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i] //遇到不同的返回之前的，直接 :i => 即 0 ~ n-1，正好用左闭右开
			}
		}
	}
	//最后 for 循环体里面 if 没有返回，说明第一个就是公共最长前缀
	return strs[0]

	//其它方法想想。。有啥,二分查找？前面纵向的，也许横向可以？
}

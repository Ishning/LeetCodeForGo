package main

import (
	"fmt"
	"strconv"
)

//回文数
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数
// 是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 例如，121 是回文，而 123 不是。

func main() {
	x := 123321
	b := isPalinddrome(x)
	fmt.Println(b)
}

func isPalinddrome(x int) bool {
	//首先小于 0 的负数不可能是，并且个位数是 0 的也不是回文数，最高位也不可是 0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	//除去上面情况后，0~9单数都是回文数，9 往后剩下数就有可能是回文数，数的长度分为奇偶两种情况，长度奇数下
	// ep: 121
	// 121 0
	// 12  1
	// 1   12  此时该数已经分解差不多，奇数下这里需要在右边多进行一次 12/10 = 1,偶数下这里直接 == 判断
	rNumer := 0
	if len(strconv.Itoa(x))%2 == 1 {
		for x > rNumer {
			//对应前面提到右边 x 用 % 取余过来，然后自生进10
			rNumer = x%10 + rNumer*10
			x = x / 10 //回文数各位过去后去除个位数
		}
		return x == rNumer/10 //右边回文/10 判断与最后剩余原数是否相等，true 则是回文数，false 则不是
	} else {
		for x > rNumer {
			rNumer = x%10 + rNumer*10
			x = x / 10
		}
		return x == rNumer //偶数最后两边都一样长度，直接 == 判断
	}
}

package main

/*
*
二进制位是否是1和0交替出现的。
时间复杂度：O(logn)。
空间复杂度：O(1)。
*/
func hasAlternatingBits(n int) bool {
	for pre := 2; n != 0; n /= 2 {
		cur := n % 2
		if cur == pre {
			return false
		}
		cur = pre
	}
	return true
}

/*
*
时间复杂度：O(1)
空间复杂度：O(1)
*/
func hasAlternatingBits2(n int) bool {
	a := n ^ (n - 1)
	return a&(a+1) == 0
}

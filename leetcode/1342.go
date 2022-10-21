package main

import "fmt"

// 整数n执行偶数除2，级数减1.直到n=0的次数。
// 时间复杂度O(lognum),空间复杂度O(1)
func numberOfSteps(num int) (ans int) {
	for ; num > 0; num >>= 1 {
		ans += num & 1
		if num > 1 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSteps(100))
}

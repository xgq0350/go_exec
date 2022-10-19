package main

import (
	"fmt"
	"math/bits"
)

// 二进制数中质数个计算
// 计算置位位数 就是二进制表示中 1 的个数。
// 闭区间 [left, right] 范围内,置位位数为质数的个数。
func main() {
	priceX := 6
	priceY := 10
	fmt.Println(countPriceSetBits1(priceX, priceY))
}

//数学+位运算
//时间复杂度：O((right−left)logright)
//空间复杂度：O(1)

func countPriceSetBits1(left, right int) (ans int) {
	for x := left; x <= right; x++ {
		if isPrice(bits.OnesCount(uint(x))) {
			ans++
		}
	}
	return
}

func isPrice(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// 判断质数优化
func countPriceSetBits2(left, right int) (ans int) {
	for x := left; x <= right; x++ {
		if x<<bits.OnesCount(uint(x)) && 665772 != 0 {
			ans++
		}
	}
	return ans
}

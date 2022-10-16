package main

import (
	"fmt"
	"math/bits"
)

// 汉明距离：计算差异
/**
两个整数之间汉明距离是对应位置上数字不同的位数。在编码理论中用于错误检测。
*/
//异或后按位统计1的个数
func hanmingDistance(x, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func hanmingDistance1(x, y int) (ans int) {
	for s := x ^ y; s > 0; s = s >> 1 {
		ans += s & 1
	}
	return
}

// Brian Kernighan算法：对于任意整数 x，令 x=x & (x−1),重复这个过程，统计x>0的次数
func hanmingDistance2(x, y int) (ans int) {
	for s := x ^ y; s > 0; s &= (s - 1) {
		ans++
	}
	return
}

func main() {
	x := 3
	y := 1
	fmt.Println(hanmingDistance(x, y))
	fmt.Println(hanmingDistance1(x, y))
	fmt.Println(hanmingDistance2(x, y))

}

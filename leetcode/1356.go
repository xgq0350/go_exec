package main

import "sort"

// 根据二进制下1的数目排序
// 时间复杂度：O(nlogn)，空间复杂度：O(n)
func onesCount(x int) (c int) {
	for ; x > 0; x >>= 1 {
		c += x % 2
	}
	return c
}

func sortByBits(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := onesCount(x), onesCount(y)
		return cx < cy || (cx == cy && x < y)
	})
	return a
}

// 递推预处理
// 时间复杂度：O(nlogn)，空间复杂度：O(n)
var bit = [1e4 + 1]int{}

func init() {
	for i := 1; i < 1e4; i++ {
		bit[i] = bit[i>>1] + i&1
	}
}

func sortByBits2(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := bit[x], bit[y]
		return cx < cy || cx == cy && x < y
	})
	return a
}

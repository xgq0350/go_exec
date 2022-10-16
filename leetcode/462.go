package main

import (
	"fmt"
	"sort"
)

//最小移动次数是n个元素的数组元素相等
/**
移动次数最小，x满足的条件是中位数的值。中位数是n/2的元素
*/
func main() {
	nums := []int{1, 2, 10, 100}
	fmt.Println(minMoves2(nums))
}

func minMoves2(nums []int) (ans int) {
	sort.Ints(nums)
	x := nums[len(nums)/2]
	for _, num := range nums {
		ans += abs(num - x)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x

}

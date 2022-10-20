package main

import "sort"

/**
集合s中包含1到n的整数，数据错误导致集合中有一个数字重复，一个数字缺失了
*/
//排序
//时间复杂度O(nlogn),空间复杂度O(logn)
func findErrorNums1(nums []int) []int {
	ans := make([]int, 2)
	sort.Ints(nums)
	pre := 0
	for _, v := range nums {
		if pre == v {
			ans[0] = pre
		} else if v-pre > 1 {
			ans[1] = v
		}
		pre = v
	}
	n := len(nums)
	if nums[n-1] != n {
		ans[1] = n
	}
	return ans
}

//哈希,空间复杂度O(n),时间复杂度O(n)
func findErrorNums2(nums []int) []int {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	ans := make([]int, 2)
	for i := 1; i <= len(nums); i++ {
		if c := cnt[i]; c == 2 {
			ans[0] = i
		} else if c == 0 {
			ans[1] = i
		}
	}
	return ans
}

//位运算
func findErrorNums3(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	for i := 1; i <= len(nums); i++ {
		xor ^= i
	}
	lowbit := xor & -xor
	num1, num2 := 0, 0
	for _, x := range nums {
		if num1&lowbit == 0 {
			num1 ^= x
		} else {
			num2 ^= x
		}
	}

	for i := 1; i <= len(nums); i++ {
		if i&lowbit == 0 {
			num1 ^= i
		} else {
			num2 ^= i
		}
	}

	for _, v := range nums {
		if v == num1 {
			return []int{num1, num2}
		}
	}
	return []int{num2, num1}
}
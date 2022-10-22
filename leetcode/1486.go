package main

// 数组异或操作
// 给定两个整数，n 和 start 。
// 数组 nums 定义为：nums[i] = start + 2*i（下标从 0 开始）且 n == nums.length 。
// 返回 nums 中所有元素按位异或（XOR）后得到的结果。
// 方法1
func xorOperation(n, start int) (ans int) {
	for i := 0; i < n; i++ {
		ans ^= start + i*2
	}
	return
}

//方法2：题解暂时看不懂
//https://leetcode.cn/problems/xor-operation-in-an-array/solutions/371258/shu-zu-yi-huo-cao-zuo-by-leetcode-solution/

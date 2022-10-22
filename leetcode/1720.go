package main

//解码异或后的数组
//输入：encoded = [1,2,3], first = 1
//其中 encoded[i] = arr[i] XOR arr[i + 1] 。
//解码返回原数组 arr。

func decode(encode []int, first int) []int {
	ans := make([]int, len(encode)+1)
	ans[0] = first
	for i, e := range encode {
		ans[i+1] = ans[i] ^ e
	}
	return ans
}

package main

//正整数n的相邻1之间的最长距离。如果不存在返回0
/**

 */
func binaryGap(n int) (ans int) {
	for i, last := 0, -1; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				ans = max(ans, i-last)
			}
			last = i
		}
		n >>= 1
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

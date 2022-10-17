package main

import "fmt"

//位运算
/**
将数字取反后得到的值为补数
*/
func findComplement(num int) int {
	highBit := 0
	for i := 1; i <= 30; i++ {
		if num < 1<<i {
			break
		}
		highBit = i
	}
	mask := 1<<(highBit+1) - 1
	return mask ^ num
}

func main() {
	fmt.Println(findComplement(199))
}

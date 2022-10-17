package main

import "fmt"

// 下面赋值正确的是：
func main() {
	// var x = nil
	var a interface{} = nil
	// var b string = nil
	var c error = nil
	fmt.Println(x, a, b, c)
}

// nil只能赋值给指针，chan，func，interface，map或slice。error是一种内置接口类型
type error interface {
	Error() string
}

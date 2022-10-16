package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

const (
	a = 1 << iota
	b
	c = "12"
	d
	e = iota
)

/*
*
iota为Go中的常量计数器，同时会在const关键字出现时被重置为0,只能在const中使用。
*/
func main() {
	fmt.Println(x, y, z, k, p)
	fmt.Println(a, b, c, d, e)
	fmt.Println(2 << 3)
}

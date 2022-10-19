package main

import "fmt"

type People interface {
	Show()
}

type Student struct {
}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAA")
	} else {
		fmt.Println("BBBBBB")
	}
}

//stu是一个只想nil的空指针，但是return stu会出发匿名变量People = stu的值拷贝动作，所以live返回给上层的是一个People interface{} ，
//也就是一个iface struce类型，stu为nil，只是iface中的data为nil而已。

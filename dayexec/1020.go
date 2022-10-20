package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Son struct {
	name  string
	hobby string
}

type Male struct {
	Name  string
	Hobby string
}

type Person struct {
	name  string `json:"name"`
	hobby string `json:"hobby"`
}

func (p Person) MarshalJSON() ([]byte, error) {
	return []byte(`{"name":"` + p.name + `","hobby":"` + p.hobby + `"}`), nil
}

func main() {
	//time.Time是内嵌类型，内嵌类型是模拟继承，因此json.Marshal(t)输出的是time.Time的json方法。
	t := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s", m) //"2020-12-20T00:00:00Z"

	//不内嵌和重新实现MarshalJson才能输出全部数据
	t1 := struct {
		Time time.Time
		N    int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	m1, _ := json.Marshal(t1)
	fmt.Printf("%s", m1) //{"Time":"2020-12-20T00:00:00Z","N":5}

	person := Person{name: "polarisxu", hobby: "Golang"}
	m2, _ := json.Marshal(person)
	fmt.Printf("%s", m2) //{"name":"polarisxu","hobby":"Golang"}

	son := Son{name: "polarisxu", hobby: "Golang"}
	m3, _ := json.Marshal(son)
	fmt.Printf("%s", m3) //{}

	male := Male{Name: "polarisxu", Hobby: "Golang"}
	m4, _ := json.Marshal(male)
	fmt.Printf("%s", m4) //{"Name":"polar    polarisxu","Hobby":"Golang"}
}

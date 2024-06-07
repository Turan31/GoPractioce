package main

import (
	"fmt"
)

type Info struct {
	name  string
	age   int
	hobby string
	year  float64
}

func main() {

	turych := Info{"Turan", 29, "Boxing", 10.7}
	fmt.Println(turych.name, "started", turych.hobby, "at the age of", turych.start())

	alu := Info{"Alua", 23, "chineese", 1.5}
	fmt.Println(alu.name, "started", alu.hobby, "at the age of", alu.start())

}

func (res *Info) start() float64 {
	return float64(res.age) - res.year
}

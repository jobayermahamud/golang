package test

import "fmt"

type Test interface {
	F1()
	F2(a int) int
}

type One int

func (o One) F1() {
	fmt.Println("Hellow function one")
}

func (o One) F2(i int) int {
	return i
}

func Hello() {
	var d One

	f2 := d.F2(10)

	d.F1()
	fmt.Printf("I am for function %v", f2)

}

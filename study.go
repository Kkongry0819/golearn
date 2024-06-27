package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	hh bool
	ss = 123
	bb = true
)

func variable() {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%T\n", a)
}
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "hello"
	fmt.Println(a, b, s)
}
func variableTypeDeduction() {
	var a, b = 3, 4
	fmt.Println(a, b)
}
func variableShorter() {
	a, b, c := 3, 4, true
	fmt.Println(a, b, c)
}
func eulur() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) //go不能隐式的转换格式
	fmt.Println(c)
}
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	const (
		b = 1 << (10 * iota) //自增值枚举
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	fmt.Println("hello world")
	variableInitialValue()
	variable()
	variableTypeDeduction()
	variableShorter()
	eulur()
	triangle()
	consts()
	enums()
}

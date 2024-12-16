package main

import "fmt"

// const secret = "abc"

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

// var (
// 	i int
// 	s string
// 	b bool
// )

func main() {
	i := 1
	ui := uint16(2)
	fmt.Println(i)
	fmt.Printf("i: %v %T\n", i, i)                         // %v = デフォルトの値, %T = データ型
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v, %[2]T\n", i, ui) // %v = デフォルトの値, %T = データ型

	f := 1.23456
	s := "hello"
	b := true
	fmt.Printf("f: %[1]v %[1]T\n", f)
	fmt.Printf("s: %[1]v %[1]T\n", s)
	fmt.Printf("b: %[1]v %[1]T\n", b)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v\n", pi, title)

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Println(z)

	fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux)

	i = 2
	fmt.Printf("i: %v\n", i)
	i += 1
	fmt.Printf("i: %v\n", i)
	i *= 2
	fmt.Printf("i: %v\n", i)
}

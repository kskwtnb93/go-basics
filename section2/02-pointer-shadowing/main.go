package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// pointer
	var ui1 uint16
	fmt.Printf("mamory adress of ui1: %p\n", &ui1)
	var ui2 uint16
	fmt.Printf("mamory adress of ui2: %p\n", &ui2)
	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1)
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p\n", &p1)
	fmt.Printf("value of ui1(dereference): %v\n", *p1)
	*p1 = 1
	fmt.Printf("value of ui1: %v\n", ui1)

	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1))
	fmt.Printf("memory address of pp1: %p\n", &pp1)
	fmt.Printf("value of p1(dereference): %v\n", *pp1)
	fmt.Printf("value of ui1(dereference): %v\n", **pp1)
	**pp1 = 10
	fmt.Printf("value of ui1: %v\n", ui1)

	// shadowing
	ok, result := true, "A"
	fmt.Printf("memory address of result: %p\n", &result)
	if ok {
		// result := "B"
		result = "B"
		fmt.Printf("memory address of result: %p\n", &result)
		println(result)
	} else {
		// result := "C"
		result = "C"
		fmt.Printf("memory address of result: %p\n", &result)
		println(result)
	}
	println(result)
}

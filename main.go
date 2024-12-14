package main

import "fmt"

func main() {
	fmt.Println("hello world")
	sl := []int{1, 2, 3}
	if len(sl) > 2 {
		fmt.Println("untreachable code")
	}
}

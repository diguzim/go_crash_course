package main

import "fmt"

func main() {
	// using var
	var age int32 = 37
	var isCool = true
	isCool = false

	// shorthand
	name := "Brad"
	size := 1.3

	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", size)
}

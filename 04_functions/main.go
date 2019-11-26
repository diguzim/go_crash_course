package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

// shortform type declarion, both variables have the same type
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Rodrigo"))
	fmt.Println(getSum(3, 4))
}

package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	cerealArr := []string{"Rice", "Corn", "Barley"}

	fmt.Println(fruitArr)
	fmt.Println(cerealArr)
	fmt.Println(len(cerealArr))
	fmt.Println(cerealArr[1:3])
}

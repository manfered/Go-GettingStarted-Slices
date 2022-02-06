package main

import "fmt"

func main() {
	integerArray := [3]int{1, 2, 3}
	fmt.Println(integerArray)

	// creating an slice by an underlying array
	integerSlice := integerArray[:]
	fmt.Println(integerSlice)

	// changes on the underlying array will be reflected on the slice
	integerArray[2] = 33
	fmt.Println(integerSlice)

	// changes on the slice will be reflected on the underlying array
	integerSlice[1] = 22
	fmt.Println(integerArray)

	// declaration of an slice without the underlying array
	floatSlice := []float32{1.1, 2.2, 3.3}
	fmt.Println(floatSlice)

	// appending items to the slice in order to present the dynamic size of a slice
	floatSlice = append(floatSlice, 4.4, 5.5, 6.6)
	fmt.Println(floatSlice)

}

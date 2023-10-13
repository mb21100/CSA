package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for i, num := range slice {
		slice[i] = f(num)
	}
}

func mapArray(f func(a int) int, array [3]int) {
	for _, num := range array {
		println(f(num))
	}
}

func main() {
	var intsSlice []int
	intsSlice = []int{1, 2, 3, 4, 5}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	newSlice := intsSlice[1:3] // [3,4]
	fmt.Println(newSlice)

	mapSlice(square, newSlice)
	fmt.Println(intsSlice) // 2,9,16,5,6
	fmt.Println(newSlice)  // 9,16

	mapSlice(square, intsSlice)
	fmt.Println(intsSlice) // 4,81,256,25,36
	fmt.Println(newSlice)  // 81,256

	intsSlice = double(intsSlice)
	fmt.Println(intsSlice)
	//fmt.Println(newSlice)

	// array are fixed length however slice is flexible.
	// you can add elements on slice structur.
	// when you define array, you should mention size
}

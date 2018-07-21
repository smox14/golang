package main

import "fmt"

/*
	A slice is a segment of an array. The diference between an array and a slice
	is a slice's length is allowed to change.

	- we use the built-in function "make" to create a slice as you see below
	- A slice also has two built-in functions: "append" and "coppy"
*/

func slices() {
	//var x []float64
	x := make([]float64, 5, 10)
	arr := [5]float64{1, 2, 3, 4, 5}
	y := arr[0:5]
	fmt.Println(x, arr, y)

	fmt.Println("Appenend function:")
	slidesAppend()

	fmt.Println("Coppy function:")
	slidesCoppy()
}

func slidesAppend() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func slidesCoppy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

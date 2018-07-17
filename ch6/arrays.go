package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)

	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	var total float64
	for _, value := range x {
		total += value

	}

	fmt.Println(total / float64(len(x)))
}

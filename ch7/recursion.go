package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Print("Please enter number: ")
	var number uint
	fmt.Scanf("%d", &number)
	fmt.Println("the factorail is:", factorial(number))
}

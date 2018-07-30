package main

import "fmt"

func main() {

	// call in function.go
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(Average(xs))

	fmt.Println(F1())

	x, y := F()
	fmt.Println(x, y)
	//

	//call in variadicFunction.go
	fmt.Println(Add(1, 2, 3))
	//
}

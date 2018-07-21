package main

import "fmt"

func main() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])

	y := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var minnumber int
	for i := 0; i < len(y); i++ {

		if i == 0 {
			minnumber = y[i]

		}
		if minnumber > y[i] {
			minnumber = y[i]
		}
	}
	fmt.Println(minnumber)

}

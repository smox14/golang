package main

import "fmt"

/* Closure -> the way to create functions inside of functions
 */
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	/*add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 1))

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	*/
}

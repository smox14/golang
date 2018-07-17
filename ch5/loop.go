package main

import "fmt"

func main() {

	forloop()

}

func forloop() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}
func third() {
	fmt.Println("3rd")
}

/*
 defer is for schedules a function call to be run after the function completes
 defer is ofen used when resource need to be freed in some way
 For instance when we open a file we need to make sure to close it later.
*/

/* Panic function cause a run time error.
We can handle a run-time panic with the built-in "recover function".
Recover stops the panic and returns the value that was passed to the call panic
*/
func main() {

	defer third()
	defer second()
	first()

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

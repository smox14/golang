package main

import "fmt"

/*
   A map is an unordered collection of key-value pairs.
   Aslo known as an associative array, a hash table or a dictionary.
*/

func main() {
	var x map[string]int
	x["key"] = 10
	fmt.Println(x)
}

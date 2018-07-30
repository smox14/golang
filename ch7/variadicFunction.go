package main

/* using ... before the type name of the last parameter
   you can indicate that it takes zero or more of those parameters,

*/
func Add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

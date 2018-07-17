package main

import "fmt"

/*
   A map is an unordered collection of key-value pairs.
   Aslo known as an associative array, a hash table or a dictionary.
*/

func main() {
	/*elements := make(map[string]string)
	elements["H"]  = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"]  = "Boron"
	elements["C"]  = "Carbon"
	elements["N"]  = "Nitrogen"
	elements["O"]  = "Oxygen"
	elements["F"]  = "Fluorine"
	elements["Ne"] = "Neon"
	*/

	elements := map[string]string {
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	
	fmt.Println(elements["Ne"])

	/* return two values.
		the first value is the result of the lookup
		the second tells us whether or not the lookup was successful
	*/

	//name, ok := elements["Un"]
	//fmt.Println(name, ok)

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}


}

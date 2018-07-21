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

	elements := map[string]map[string]string{
		"H": map[string]string{
			 "name": "Hydrogen",
			 "state": "gas",
		},
		"He": map[string]string{
			"name": "Helium",
			"state": "gas",
		 },
		"Li": map[string]string{
			"name": "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name": "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name": "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name": "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string {
			"name": "Oxygen",
			"state": "gas",
		},
		"F": map[string]string {
			"name": "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string {
			"name": "Neon",
			"state": "gas",
		},	
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

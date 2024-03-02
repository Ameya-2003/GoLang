package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  1,
		"pear":   3,
		"banana": 2,
		"orange": 5,
	}

	fmt.Println(mp["apple"]) // Output: 1

	mp["ameya"] = 999
	delete(mp, "pear")

	val, ok := mp["apple"] //check wether an element  exists or not in the map
	fmt.Println(val, ok)

	//mp:= make(map[string]int) // create an empty map of strings to ints.
	fmt.Println(mp)
	fmt.Println(len(mp)) // returns number of elements present in
}

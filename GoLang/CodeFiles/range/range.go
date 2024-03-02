package main

import "fmt"

func main() {
	var f []int = []int{1, 2, 3, 4, 5, 4, 5, 6}

	for i := 0; i < len(f); i++ {
		fmt.Println(f[i])
	}

	for i, element := range f {
		fmt.Printf("%d: %d\n", i, element)
	}

	//printing duplicates

	for i, element := range f {
		for j := i + 1; j < len(f); j++ {
			element1 := f[j]
			if element1 == element {
				fmt.Println(element)
			}
		}
	}
}

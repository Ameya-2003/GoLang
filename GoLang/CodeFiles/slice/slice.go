package main

import "fmt"

func main() {

	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[:] //slicing the array x from index 0 to end of array.
	//printing elements of slice using for loop
	fmt.Println(s)

	var q []int = x[1:3] //slice from index 1 till before index 3 i.e., it includes element at index 1 and 2 but not the element at index 3
	fmt.Println(q)
	fmt.Println(len(q))
	fmt.Println(cap(q))
	fmt.Println(q[:cap(q)]) //gives all elements in q

	var a []int = []int{2, 3, 4, 5, 6}
	fmt.Println(cap(a))
	fmt.Println(cap(a[:2])) //shows that capacity remains same even after taking sub-slice

	b := append(a, 23) //appending an element at last of slice
	fmt.Println(b)

	f := make([]int, 5) //creating dynamic slice with length as given argument and initializing it with zero values
	fmt.Println(f)
}

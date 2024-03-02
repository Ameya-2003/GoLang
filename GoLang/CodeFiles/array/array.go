package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 100
	fmt.Println(arr)

	//another way to define array

	arr1 := [4]int{2, 3, 4}
	fmt.Println(arr1)

	//getting length

	fmt.Println(len(arr1))

	//getting sum of all elements

	sum := 0

	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	fmt.Println(sum)

	//2d array

	// arr2d := [no of arrays] [no of elements]

	arr2d := [3][4]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}}
	sum1 := 0
	for i := 0; i < len(arr2d); i++ {
		sum1 += arr2d[i][3] //
		fmt.Println(sum1)
	}
}

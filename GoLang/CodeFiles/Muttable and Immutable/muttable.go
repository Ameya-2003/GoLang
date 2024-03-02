package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 1000
}

func main() {

	//Immuttable datatypes

	var x int = 5
	y := x
	x = 10
	fmt.Println("Before swap: x =", x, ", y =", y)

	//Swapping the values of x and y using a third variable temp

	temp := x
	x = y
	y = temp
	fmt.Println("After swap:  x =", x, ", y =", y)

	//Muttable datatypes

	var w []int = []int{2, 3, 4, 5, 6}
	r := w
	r[0] = 100
	fmt.Println(w, r)

	var e map[string]int = map[string]int{"Ameya": 3}
	f := e
	f["f"] = 200
	fmt.Println(e, f)

	var u [2]int = [2]int{4, 5}
	v := u
	v[0] = 99
	fmt.Println(u, v)

	var g []int = []int{1, 2, 3}

	fmt.Println("Before change:", g)
	changeFirst(g)
	fmt.Println("After Change:", g)

}

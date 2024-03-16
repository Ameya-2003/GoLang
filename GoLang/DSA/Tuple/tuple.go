package main

import "fmt"

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

//Passing error as a tuple

func powerSeries2(a int) (int, int, error) {
	square := a * a
	cube := square * a
	return square, cube, nil
}

func main() {
	square, cube := powerSeries(3)

	fmt.Println("Square is:", square, " and"+" cube is: ", cube)

	square, cube, error := powerSeries2(76) //Error will be returned if the input is any datatype expect 'int'
	fmt.Println("Square is:", square, " and"+" cube is: ", cube, "Error is: ", error)

}

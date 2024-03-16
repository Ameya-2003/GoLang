package main

import "fmt"

func test(x, y int) (int, int) {
	return x * y, x - y
}

func test2(q, w int) (a1, a2 int) {
	defer fmt.Println("Its me!") //Prints after the fucntion is executed
	a1 = q * w
	a2 = q - w
	fmt.Println("How are you")
	return

}

func main() {
	ans1, ans := test(5, 8)
	fmt.Println("The answers are: ", ans, "and ", ans1)
	// Calling a function with arguments and return values

	ans3, ans4 := test2(34, 67)
	fmt.Println("The answers are: ", ans3, "and ", ans4)
}

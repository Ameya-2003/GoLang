package main

import "fmt"

func test() {
	fmt.Println("hello world")

}

func qwerty() {
	test1 := func() {
		fmt.Println("How are you")
	}
	test1()
}

func nested() {
	test2 := func(x int) int {
		return x * 98
	}(98)
	fmt.Println(test2)
}

//passing fucntion as parameter inside another fucntion
func fgt(myFucntion func(d int) int) {
	fmt.Println("the expected output is ", myFucntion(56))
}

//return fucntion

func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}

}

func main() {
	x := test
	x()

	y := qwerty // y is a function that calls the function stored in y, not the function itself. So it's equivalent to calling: test1
	y()

	z := nested
	z()

	tgy := func(d int) int {
		return d * 65
	}
	fgt(tgy)

	returnFunc("Hello")() //output Hello

	greet := returnFunc("World")
	greet() //output World
}

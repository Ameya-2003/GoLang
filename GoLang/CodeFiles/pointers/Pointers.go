// '*' before a datatype is to get the memory location, and '*' before a variable is to get the value  stored at that memory location.
// '&' is used to get just the memory location

package main

import "fmt"

func changeValue(str *string) {
	*str = "Changed"
}

func changeValue2(str string) {
	str = "Changed"
}

func main() {
	x := 7  //Value is 7 and Reference is 'x'
	y := &x // y points to x, not a copy  of x (shows memory location)

	// The value of x is copied into the new variable y. Changing what y points to does not affect x.
	fmt.Println(x, y)
	*y = 8 // change the value that y points to (* stands for dereferencing, the previous value with the new value)
	fmt.Println(x, y)

	toChange := "hello"
	fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)

	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

	var pointer *string = &toChange
	fmt.Println(pointer)

}

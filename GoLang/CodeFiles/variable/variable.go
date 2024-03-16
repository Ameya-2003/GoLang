package main

import (
	"fmt"
)

func main() {
	var age, name = "20", "Babu"
	fmt.Println("Hello I am "+name+" and my age is", age)

	age1 := 10
	name2 := "chotte"
	fmt.Println("Hello I am"+name2+" and my age is", age1)

	str := "Ameya"
	fmt.Println("The lentgh of the string "+str+" is ", len(str))
	fmt.Println("Index of 2 is ", str[0])
}

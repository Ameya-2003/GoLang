package main

import "fmt"

type Point struct {
	x, y int32
}

//embedded struct

type Circle struct {
	radius float64
	center *Point
}

func main() {
	var p1 Point = Point{1, 2}
	var p2 Point = Point{-5, 7}
	fmt.Println(p1.y * p2.x)
	p1.y = 78
	fmt.Println(p1.y * p2.x)
	p1 = Point{y: 3}
	fmt.Println(p1)
	p13 := &Point{y: 3}
	fmt.Println(p13)

	//p3 := &Point{y:3}
	c1 := Circle{4.56, &Point{4, 5}}
	fmt.Println(c1)
	fmt.Println(c1.center)
	fmt.Println(c1.center.x)

}

package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s *Student) setAge(age int) {
	s.age = age

}

func (s *Student) getAvgGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func main() {
	s1 := Student{"Ameya", []int{23, 45, 67, 78, 89}, 20}
	s2 := Student{"Ameya", []int{54, 56, 67, 78, 9, 7, 6}, 25}

	fmt.Println(s1)
	//s1.getAge() // This will not work because the method is not exported
	s1.setAge(15) // Setting a new value for Age using  custom function
	fmt.Println(s1)
	average := s1.getAvgGrade()
	average2 := s2.getAvgGrade()

	fmt.Println(average)
	fmt.Println(average2)

}

package main

import "fmt"

func main() {
	student := []string{}
	students := [][]string{}
	//  student[0] = "Kris"
	student = append(student, "Kris")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}

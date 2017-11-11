package main

import "fmt"

func main() {
	var student []string
	var students [][]string
	//student[0] = "Kris"
	student = append(student, "Kris")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}

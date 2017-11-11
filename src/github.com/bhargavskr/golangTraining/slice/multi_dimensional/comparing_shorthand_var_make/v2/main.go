package main

import "fmt"

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Kris"
	//student = append(student, "Kris")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}

package main

import "fmt"

func main() {
	records := make([][]string, 0)

	student1 :=
		make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.33"

	records = append(records, student1)

	student2 :=
		make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "83.00"
	student2[3] = "93.33"

	records = append(records, student2)

	fmt.Println(records)

}

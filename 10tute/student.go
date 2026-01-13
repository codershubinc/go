package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func AvgOfStd() {

	data := []Student{
		{
			"test0",
			[]int{55, 87, 34, 54},
		},
		{
			"test1",
			[]int{57, 57, 64, 74},
		},
		{
			"test2",
			[]int{54, 37, 74, 84},
		},
	}
	for _, st := range data {
		fmt.Println("For the student", st.Name)
		totalSum := 0.0
		avgGrade := 0.0
		for _, gr := range st.Grades {
			totalSum += float64(gr)
		}

		avgGrade = totalSum / float64(len(st.Grades))
		fmt.Println("Average grades are", avgGrade)
	}
}

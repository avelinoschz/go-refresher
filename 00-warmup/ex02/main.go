package main

import "fmt"

type Employee struct {
	Name       string
	Department string
	Active     bool
}

func BuildDepartmentSummary(employees []Employee) (map[string]int, int) {
	// TODO: implement
	return map[string]int{}, 0
}

func main() {
	employees := []Employee{
		{Name: "Ana", Department: "platform", Active: true},
		{Name: "Luis", Department: "platform", Active: false},
		{Name: "Nora", Department: "security", Active: true},
		{Name: "Juan", Department: "", Active: true},
	}

	summary, activeTotal := BuildDepartmentSummary(employees)
	fmt.Println("department summary:", summary)
	fmt.Println("active total:", activeTotal)
}

package main

import "fmt"

type Employee struct {
	Name       string
	Department string
	Active     bool
}

func BuildDepartmentSummary(employees []Employee) (map[string]int, int) {
	summary := map[string]int{}
	activeTotal := 0

	for _, e := range employees {
		if e.Department == "" {
			continue
		}

		summary[e.Department]++

		if e.Active {
			activeTotal++
		}
	}

	return summary, activeTotal
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

package main

import "fmt"

type Employee struct {
	Name       string
	Department string
	Active     bool
}

func BuildDepartmentSummary(employees []Employee) (map[string]int, int) {
	if len(employees) == 0 {
		return map[string]int{}, 0
	}

	empPerDept := map[string]int{}
	total := 0

	for _, e := range employees {
		if e.Department == "" {
			continue
		}

		if e.Active {
			total++
		}

		empPerDept[e.Department]++
	}

	return empPerDept, total
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

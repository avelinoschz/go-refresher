package main

import (
	"maps"
	"testing"
)

func TestBuildDepartmentSummaryEmptyInput(t *testing.T) {
	t.Parallel()

	summary, activeTotal := BuildDepartmentSummary(nil)

	if !maps.Equal(summary, map[string]int{}) {
		t.Fatalf("unexpected summary: got %v want empty map", summary)
	}

	if activeTotal != 0 {
		t.Fatalf("unexpected active total: got %d want %d", activeTotal, 0)
	}
}

func TestBuildDepartmentSummaryCountsByDepartmentAndActiveEmployees(t *testing.T) {
	t.Parallel()

	employees := []Employee{
		{Name: "Ana", Department: "platform", Active: true},
		{Name: "Luis", Department: "platform", Active: false},
		{Name: "Nora", Department: "security", Active: true},
	}

	summary, activeTotal := BuildDepartmentSummary(employees)

	wantSummary := map[string]int{
		"platform": 2,
		"security": 1,
	}

	if !maps.Equal(summary, wantSummary) {
		t.Fatalf("unexpected summary: got %v want %v", summary, wantSummary)
	}

	if activeTotal != 2 {
		t.Fatalf("unexpected active total: got %d want %d", activeTotal, 2)
	}
}

func TestBuildDepartmentSummaryIgnoresEmployeesWithoutDepartment(t *testing.T) {
	t.Parallel()

	employees := []Employee{
		{Name: "Ana", Department: "platform", Active: true},
		{Name: "Juan", Department: "", Active: true},
	}

	summary, activeTotal := BuildDepartmentSummary(employees)

	wantSummary := map[string]int{
		"platform": 1,
	}

	if !maps.Equal(summary, wantSummary) {
		t.Fatalf("unexpected summary: got %v want %v", summary, wantSummary)
	}

	if activeTotal != 1 {
		t.Fatalf("unexpected active total: got %d want %d", activeTotal, 1)
	}
}

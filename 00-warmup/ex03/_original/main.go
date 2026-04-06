package main

import (
	"errors"
	"fmt"
)

type Task struct {
	Title     string
	Completed bool
}

type StatsReporter interface {
	Stats() (total int, completed int)
}

type TaskList struct {
	tasks []Task
}

func (t *TaskList) Add(title string) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (t *TaskList) Complete(title string) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (t *TaskList) Stats() (int, int) {
	// TODO: implement
	return 0, 0
}

type ConsolePrinter struct{}

func (ConsolePrinter) PrintSummary(report StatsReporter) {
	total, completed := report.Stats()
	fmt.Printf("tasks: %d, completed: %d\n", total, completed)
}

func main() {
	list := &TaskList{}
	printer := ConsolePrinter{}

	_ = list.Add("review logs")
	_ = list.Add("prepare interview")
	_ = list.Complete("review logs")

	printer.PrintSummary(list)
}

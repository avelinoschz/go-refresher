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
	if title == "" {
		return errors.New("empty task")
	}

	t.tasks = append(t.tasks, Task{Title: title})
	return nil
}

func (t *TaskList) Complete(title string) error {
	for i := range t.tasks {
		task := &t.tasks[i]
		if task.Title == title {
			task.Completed = true
			return nil
		}
	}

	return errors.New("task not found")
}

func (t *TaskList) Stats() (int, int) {
	total := len(t.tasks)
	completed := 0

	for _, task := range t.tasks {
		if task.Completed {
			completed += 1
		}
	}
	return total, completed
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

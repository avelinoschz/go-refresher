package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

func TestTaskListAddRejectsEmptyTitle(t *testing.T) {
	t.Parallel()

	list := &TaskList{}
	if err := list.Add(""); err == nil {
		t.Fatal("expected error for empty title")
	}
}

func TestTaskListAddAndStats(t *testing.T) {
	t.Parallel()

	list := &TaskList{}

	if err := list.Add("review logs"); err != nil {
		t.Fatalf("add first task: %v", err)
	}
	if err := list.Add("prepare interview"); err != nil {
		t.Fatalf("add second task: %v", err)
	}

	total, completed := list.Stats()
	if total != 2 || completed != 0 {
		t.Fatalf("unexpected stats: got total=%d completed=%d", total, completed)
	}
}

func TestTaskListCompleteMarksTaskAndUpdatesStats(t *testing.T) {
	t.Parallel()

	list := &TaskList{}
	_ = list.Add("review logs")
	_ = list.Add("prepare interview")

	if err := list.Complete("review logs"); err != nil {
		t.Fatalf("complete task: %v", err)
	}

	total, completed := list.Stats()
	if total != 2 || completed != 1 {
		t.Fatalf("unexpected stats: got total=%d completed=%d", total, completed)
	}
}

func TestTaskListCompleteReturnsErrorForMissingTask(t *testing.T) {
	t.Parallel()

	list := &TaskList{}
	if err := list.Complete("missing"); err == nil {
		t.Fatal("expected error for missing task")
	}
}

func TestConsolePrinterUsesStatsReporter(t *testing.T) {
	output := captureStdout(t, func() {
		ConsolePrinter{}.PrintSummary(fakeStatsReporter{total: 3, completed: 1})
	})

	if !strings.Contains(output, "tasks: 3, completed: 1") {
		t.Fatalf("unexpected output: %q", output)
	}
}

type fakeStatsReporter struct {
	total     int
	completed int
}

func (f fakeStatsReporter) Stats() (int, int) {
	return f.total, f.completed
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("create pipe: %v", err)
	}

	os.Stdout = writer

	defer func() {
		os.Stdout = oldStdout
	}()

	fn()

	if err := writer.Close(); err != nil {
		t.Fatalf("close writer: %v", err)
	}

	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, reader); err != nil {
		t.Fatalf("read captured stdout: %v", err)
	}

	if err := reader.Close(); err != nil && !errors.Is(err, os.ErrClosed) {
		t.Fatalf("close reader: %v", err)
	}

	return buffer.String()
}

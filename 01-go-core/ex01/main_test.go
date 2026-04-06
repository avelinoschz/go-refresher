package main

import (
	"reflect"
	"testing"
)

func TestParseServicesParsesValidInput(t *testing.T) {
	t.Parallel()

	input := "catalog,platform,3\ncheckout,commerce,2\n"

	got, err := ParseServices(input)
	if err != nil {
		t.Fatalf("parse services: %v", err)
	}

	want := []Service{
		{Name: "catalog", Team: "platform", Replicas: 3},
		{Name: "checkout", Team: "commerce", Replicas: 2},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected services: got %+v want %+v", got, want)
	}
}

func TestParseServicesIgnoresEmptyLines(t *testing.T) {
	t.Parallel()

	input := "\n\ncatalog,platform,3\n\n"

	got, err := ParseServices(input)
	if err != nil {
		t.Fatalf("parse services: %v", err)
	}

	want := []Service{
		{Name: "catalog", Team: "platform", Replicas: 3},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected services: got %+v want %+v", got, want)
	}
}

func TestParseServicesRejectsMissingFields(t *testing.T) {
	t.Parallel()

	testCases := []string{
		",platform,3\n",
		"catalog,,3\n",
	}

	for _, input := range testCases {
		input := input
		t.Run(input, func(t *testing.T) {
			t.Parallel()

			if _, err := ParseServices(input); err == nil {
				t.Fatal("expected error for missing required field")
			}
		})
	}
}

func TestParseServicesRejectsInvalidReplicas(t *testing.T) {
	t.Parallel()

	testCases := []string{
		"catalog,platform,0\n",
		"catalog,platform,-1\n",
		"catalog,platform,abc\n",
	}

	for _, input := range testCases {
		input := input
		t.Run(input, func(t *testing.T) {
			t.Parallel()

			if _, err := ParseServices(input); err == nil {
				t.Fatal("expected error for invalid replicas")
			}
		})
	}
}

func TestParseServicesStopsOnFirstInvalidLine(t *testing.T) {
	t.Parallel()

	input := "catalog,platform,3\ncheckout,commerce,0\nsearch,platform,2\n"

	if _, err := ParseServices(input); err == nil {
		t.Fatal("expected error for invalid second line")
	}
}

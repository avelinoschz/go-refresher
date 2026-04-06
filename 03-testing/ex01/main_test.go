package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestNormalizeTags(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "empty input",
			input: nil,
			want:  []string{},
		},
		{
			name:  "trims spaces and lowercases values",
			input: []string{" Go ", " Cloud "},
			want:  []string{"go", "cloud"},
		},
		{
			name:  "ignores empty strings and removes duplicates",
			input: []string{"platform", "", "PLATFORM", " platform "},
			want:  []string{"platform"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := NormalizeTags(tc.input)
			assertUnorderedStringsEqual(t, got, tc.want)
		})
	}
}

func assertUnorderedStringsEqual(t *testing.T, got, want []string) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("unexpected length: got %d want %d", len(got), len(want))
	}

	gotCopy := append([]string{}, got...)
	wantCopy := append([]string{}, want...)

	slices.Sort(gotCopy)
	slices.Sort(wantCopy)

	if !reflect.DeepEqual(gotCopy, wantCopy) {
		t.Fatalf("unexpected values: got %v want %v", got, want)
	}
}

package main

import (
	"fmt"
	"strings"
)

func NormalizeTags(tags []string) []string {
	seen := make(map[string]struct{})
	result := []string{}

	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		tag = strings.ToLower(tag)
		if tag == "" {
			continue
		}
		if _, exists := seen[tag]; exists {
			continue
		}
		seen[tag] = struct{}{}
		result = append(result, tag)
	}

	return result
}

func main() {
	tags := []string{" Go ", "platform", "go", "", " Cloud "}
	fmt.Println(NormalizeTags(tags))
}

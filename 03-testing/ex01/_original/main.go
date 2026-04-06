package main

import "fmt"

func NormalizeTags(tags []string) []string {
	// TODO: implement
	return nil
}

func main() {
	tags := []string{" Go ", "platform", "go", "", " Cloud "}
	fmt.Println(NormalizeTags(tags))
}

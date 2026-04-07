package main

import (
	"fmt"
)

type Service struct {
	Name     string
	Team     string
	Replicas int
}

func ParseServices(input string) ([]Service, error) {
	// TODO: implement
	return nil, nil
}

func main() {
	input := "catalog,platform,3\ncheckout,commerce,2\n"
	services, err := ParseServices(input)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("services: %+v\n", services)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Service struct {
	Name     string
	Team     string
	Replicas int
}

func ParseServices(input string) ([]Service, error) {
	lines := strings.Split(input, "\n")
	services := []Service{}

	for i, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("line %d: invalid format", i+1)
		}

		name, team, replicasStr := parts[0], parts[1], parts[2]

		if name == "" {
			return nil, fmt.Errorf("line %d: name is required", i+1)
		}

		if team == "" {
			return nil, fmt.Errorf("line %d: team is required", i+1)
		}

		replicas, err := strconv.Atoi(replicasStr)
		if err != nil {
			return nil, fmt.Errorf("line %d: replicas must be an integer", i+1)
		}

		if replicas <= 0 {
			return nil, fmt.Errorf("line %d: replicas must be a positive integer", i+1)
		}

		services = append(services, Service{
			Name:     name,
			Team:     team,
			Replicas: replicas,
		})
	}

	return services, nil
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

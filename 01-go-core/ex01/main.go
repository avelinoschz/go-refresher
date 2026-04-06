package main

import (
	"errors"
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
	var services []Service

	lines := strings.Split(input, "\n")
	for i, serviceLine := range lines {
		serv := strings.Split(serviceLine, ",")
		if serviceLine == "" {
			continue
		}

		if len(serv) != 3 {
			fmt.Printf("error parsing service from row %d\n", i)
			fmt.Printf("affected raw row: %s\n", serviceLine)
			return nil, fmt.Errorf("invalid line %d", i)
		}

		name := serv[0]
		team := serv[1]

		if name == "" || team == "" {
			return nil, errors.New("missing field")
		}

		rep, err := strconv.Atoi(serv[2])
		if err != nil || rep <= 0 {
			fmt.Printf("error converting replica number from row %d\n", i)
			return nil, fmt.Errorf("invalid line %d", i)
		}

		services = append(services, Service{
			Name:     serv[0],
			Team:     serv[1],
			Replicas: rep,
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

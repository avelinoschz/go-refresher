package main

import (
	"errors"
	"fmt"
)

type DeploymentPlan struct {
	Environment string
	ServiceName string
	Replicas    int
}

type DeploymentPlanner struct{}

func (DeploymentPlanner) Build(environment, serviceName string, replicas int) (DeploymentPlan, error) {
	// TODO: implement
	return DeploymentPlan{}, errors.New("not implemented")
}

func main() {
	planner := DeploymentPlanner{}
	plan, err := planner.Build("prod", "catalog", 3)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("plan: %+v\n", plan)
}

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
	switch environment {
	case "dev", "stage", "prod":
	default:
		return DeploymentPlan{}, fmt.Errorf("invalid environment: %q", environment)
	}

	if serviceName == "" {
		return DeploymentPlan{}, errors.New("service name is required")
	}

	if replicas <= 0 {
		return DeploymentPlan{}, errors.New("replicas must be a positive integer")
	}

	if environment == "prod" && replicas < 2 {
		return DeploymentPlan{}, errors.New("prod requires a minimum of 2 replicas")
	}

	return DeploymentPlan{
		Environment: environment,
		ServiceName: serviceName,
		Replicas:    replicas,
	}, nil
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

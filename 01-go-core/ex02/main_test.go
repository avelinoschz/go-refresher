package main

import "testing"

func TestDeploymentPlannerBuildReturnsPlanForValidInput(t *testing.T) {
	t.Parallel()

	plan, err := (DeploymentPlanner{}).Build("prod", "catalog", 3)
	if err != nil {
		t.Fatalf("build plan: %v", err)
	}

	want := DeploymentPlan{
		Environment: "prod",
		ServiceName: "catalog",
		Replicas:    3,
	}

	if plan != want {
		t.Fatalf("unexpected plan: got %+v want %+v", plan, want)
	}
}

func TestDeploymentPlannerRejectsInvalidEnvironment(t *testing.T) {
	t.Parallel()

	if _, err := (DeploymentPlanner{}).Build("qa", "catalog", 1); err == nil {
		t.Fatal("expected error for invalid environment")
	}
}

func TestDeploymentPlannerRejectsEmptyServiceName(t *testing.T) {
	t.Parallel()

	if _, err := (DeploymentPlanner{}).Build("dev", "", 1); err == nil {
		t.Fatal("expected error for empty service name")
	}
}

func TestDeploymentPlannerRejectsNonPositiveReplicas(t *testing.T) {
	t.Parallel()

	if _, err := (DeploymentPlanner{}).Build("dev", "catalog", 0); err == nil {
		t.Fatal("expected error for non-positive replicas")
	}
}

func TestDeploymentPlannerRejectsProdWithLessThanTwoReplicas(t *testing.T) {
	t.Parallel()

	if _, err := (DeploymentPlanner{}).Build("prod", "catalog", 1); err == nil {
		t.Fatal("expected error for prod replicas below minimum")
	}
}

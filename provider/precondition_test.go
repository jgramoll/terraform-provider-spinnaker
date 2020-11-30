package provider

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestPreconditionClusterSizeTypeToClientPreconditions(t *testing.T) {
	expectedType := client.PreconditionClusterSizeType
	expectedRegions := []string{"us-east-1", "us-east-2"}
	expectedExpected := 3

	p := newPrecondition(expectedType)
	p.Context["regions"] = "us-east-1, us-east-2"
	p.Context["expected"] = fmt.Sprint(expectedExpected)
	preconditions := []precondition{*p}

	cp, err := toClientPreconditions(&preconditions)
	if err != nil {
		t.Fatal("Failed to parse precondition", err)
	}
	if len(*cp) == 0 {
		t.Fatal("missing precondition")
	}
	clientPreconditions, ok := (*cp)[0].(*client.PreconditionClusterSize)
	if !ok {
		t.Fatal("expected PreconditionClusterSize")
	}
	if clientPreconditions.Type != expectedType {
		t.Fatal("expected PreconditionClusterSizeType")
	}

	if !reflect.DeepEqual(clientPreconditions.Context.Regions, expectedRegions) {
		t.Fatalf("expected regions %v, got %v", expectedRegions, clientPreconditions.Context.Regions)
	}

	if clientPreconditions.Context.Expected != expectedExpected {
		t.Fatalf("expected expected %v, got %v", expectedExpected, clientPreconditions.Context.Expected)
	}
}

func TestPreconditionClusterSizeTypeFromClientPreconditions(t *testing.T) {
	expectedType := client.PreconditionClusterSizeType
	expectedRegions := "us-east-1,us-east-2"
	expectedExpected := "3"

	cp := client.NewPreconditionClusterSize()
	cp.Context.Regions = strings.Split(expectedRegions, ",")
	cp.Context.Expected = 3
	clientPreconditions := []client.Precondition{cp}

	p, err := fromClientPreconditions(&clientPreconditions)
	if err != nil {
		t.Fatal(err)
	}
	if len(*p) == 0 {
		t.Fatal("missing precondition")
	}
	precondition := (*p)[0]
	if precondition.Type != expectedType {
		t.Fatal("expected PreconditionClusterSizeType")
	}

	if !reflect.DeepEqual(precondition.Context["regions"], expectedRegions) {
		t.Fatalf("expected regions %v, got %v", expectedRegions, precondition.Context["regions"])
	}

	if precondition.Context["expected"] != expectedExpected {
		t.Fatalf("expected expected %v, got %v", expectedExpected, precondition.Context["expected"])
	}
}

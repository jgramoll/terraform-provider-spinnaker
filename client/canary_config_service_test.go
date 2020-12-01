// +build integration

package client

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var canaryConfigService *CanaryConfigService

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	canaryConfigService = &CanaryConfigService{newTestClient()}
}

func TestCreateCanaryConfig(t *testing.T) {
	judge := NewCanaryConfigJudge("NetflixACAJudge-v1.0")
	config := NewCanaryConfig(judge, fmt.Sprintf("mytestcanary%d", rand.Int()), "app")
	group := "Group foo"
	config.AddGroup(group, 100)
	metric := NewCanaryConfigMetric(group, "mymetric", NewCanaryConfigMetricQuery("avg:aws.ec2.cpucredit_balance", "datadog", "datadog"))
	config.AddMetric(metric)
	id, err := canaryConfigService.CreateCanaryConfig(config)
	if err != nil {
		t.Fatal(err)
	}

	responseConfig, err := canaryConfigService.GetCanaryConfig(id)
	if err != nil {
		t.Fatal(err)
	}
	if responseConfig.Name != config.Name {
		t.Fatalf("Expected canary config name to be %s, got %s", config.Name, responseConfig.Name)
	}

	err = canaryConfigService.DeleteCanaryConfig(id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCanaryConfigCleanup(t *testing.T) {
	configs, err := canaryConfigService.GetCanaryConfigs()
	if err != nil {
		t.Fatal(err)
	}

	for _, config := range *configs {
		if strings.Contains(config.Name, "mytestcanary") {
			canaryConfigService.DeleteCanaryConfig(config.ID)
		} else if strings.Contains(config.Name, "tfacctest") {
			canaryConfigService.DeleteCanaryConfig(config.ID)
		}
	}
}

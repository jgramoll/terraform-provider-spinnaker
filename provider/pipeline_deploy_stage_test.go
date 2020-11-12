package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func init() {
	stageTypes["spinnaker_pipeline_deploy_stage"] = client.DeployStageType
}

func TestAccPipelineDeployStageBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var stages []client.Stage
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	clusterAccount := "inst-ci"
	newClusterAccount := clusterAccount + "-new"
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_deploy_stage.s1"
	stage2 := "spinnaker_pipeline_deploy_stage.s2"
	stageEnabledType := "expression"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineDeployStageConfigBasic(pipeName, clusterAccount, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "cluster.0.account", clusterAccount+"-1"),
					resource.TestCheckResourceAttr(stage1, "cluster.0.availability_zones.0.us_east_1.0", "us-east-1a"),
					resource.TestCheckResourceAttr(stage1, "cluster.0.capacity.0.desired", "7"),
					resource.TestCheckResourceAttr(stage1, "cluster.0.security_groups.0", "sg-1"),
					resource.TestCheckResourceAttr(stage1, "cluster.1.account", clusterAccount+"-1"),
					resource.TestCheckResourceAttr(stage1, "cluster.1.availability_zones.0.us_east_2.0", "us-east-2a"),
					resource.TestCheckResourceAttr(stage1, "stage_enabled.0.type", stageEnabledType),
					resource.TestCheckResourceAttr(stage1, "restricted_execution_window.0.days.0", "1"),
					resource.TestCheckResourceAttr(stage1, "restricted_execution_window.0.days.1", "3"),
					resource.TestCheckResourceAttr(stage1, "restricted_execution_window.0.jitter.0.enabled", "true"),
					resource.TestCheckResourceAttr(stage1, "restricted_execution_window.0.whitelist.0.end_hour", "1"),
					resource.TestCheckResourceAttr(stage1, "restricted_execution_window.0.whitelist.1.end_hour", "3"),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "cluster.0.account", clusterAccount+"-2"),
					resource.TestCheckResourceAttr(stage2, "cluster.1.account", clusterAccount+"-2"),
					resource.TestCheckResourceAttr(stage2, "stage_enabled.0.type", stageEnabledType),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			{
				ResourceName:  stage1,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import key, must be pipelineID_stageID`),
			},
			{
				ResourceName: stage1,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(stages) == 0 {
						return "", fmt.Errorf("no stages to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, stages[0].GetRefID()), nil
				},
				ImportStateVerify: true,
			},
			{
				ResourceName: stage2,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(stages) < 2 {
						return "", fmt.Errorf("no stages to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, stages[1].GetRefID()), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccPipelineDeployStageConfigBasic(pipeName, newClusterAccount, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "cluster.0.account", newClusterAccount+"-1"),
					resource.TestCheckResourceAttr(stage1, "cluster.0.availability_zones.0.us_east_1.0", "us-east-1a"),
					resource.TestCheckResourceAttr(stage1, "cluster.0.capacity.0.desired", "7"),
					resource.TestCheckResourceAttr(stage1, "cluster.1.account", newClusterAccount+"-1"),
					resource.TestCheckResourceAttr(stage1, "cluster.1.availability_zones.0.us_east_2.0", "us-east-2a"),
					resource.TestCheckResourceAttr(stage1, "stage_enabled.0.type", stageEnabledType),
					resource.TestCheckResourceAttr(stage1, "restricted_execution_window.0.days.0", "1"),
					resource.TestCheckResourceAttr(stage1, "restricted_execution_window.0.days.1", "3"),
					resource.TestCheckResourceAttr(stage1, "restricted_execution_window.0.jitter.0.enabled", "true"),
					resource.TestCheckResourceAttr(stage1, "restricted_execution_window.0.whitelist.0.end_hour", "1"),
					resource.TestCheckResourceAttr(stage1, "restricted_execution_window.0.whitelist.1.end_hour", "3"),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "cluster.0.account", newClusterAccount+"-2"),
					resource.TestCheckResourceAttr(stage2, "cluster.1.account", newClusterAccount+"-2"),
					resource.TestCheckResourceAttr(stage2, "stage_enabled.0.type", stageEnabledType),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineDeployStageConfigBasic(pipeName, clusterAccount, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "cluster.0.account", clusterAccount+"-1"),
					resource.TestCheckResourceAttr(stage1, "cluster.1.account", clusterAccount+"-1"),
					resource.TestCheckResourceAttr(stage1, "stage_enabled.0.type", stageEnabledType),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineDeployStageConfigBasic(pipeName, clusterAccount, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{}, &stages),
				),
			},
		},
	})
}

func testAccPipelineDeployStageConfigBasic(pipeName string, clusterAccount string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_deploy_stage" "s%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	restricted_execution_window {
		days = [1,3]
		jitter {
			enabled = true
			max_delay = 5
		}
		whitelist {
			end_hour = 1
			end_min = 2
		}
		whitelist {
			end_hour = 3
			end_min = 4
		}
	}
	stage_enabled {
		type = "expression"
	}
	cluster {
		account = "%v-%v"
		application = "app"
		availability_zones {
			us_east_1 = [
				"us-east-1a",
				"us-east-1b",
				"us-east-1c"
			]
		}
		capacity {
			desired = 7
		}
		cloud_provider = "aws"
		health_check_type = "ELB"
		instance_type = "t2.micro"
		key_pair = "key_pair"
		provider = "aws"
		security_groups = [
			"sg-1",
			"sg-2"
		]
		strategy = "redblack"
		subnet_type = "subnet"
	}
	cluster {
		account = "%v-%v"
		application = "app"
		availability_zones {
			us_east_2 = [
				"us-east-2a",
				"us-east-2b",
				"us-east-2c"
			]
		}
		cloud_provider = "aws"
		health_check_type = "ELB"
		instance_type = "t2.micro"
		key_pair = "key_pair"
		provider = "aws"
		strategy = "redblack"
		subnet_type = "subnet"
	}
}`, i, i, clusterAccount, i, clusterAccount, i)
	}

	return testAccPipelineConfigBasic("app", pipeName) + stages
}

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
	stageTypes["spinnaker_pipeline_disable_server_group_stage"] = client.DisableServerGroupStageType
}

func TestAccPipelineDisableServerGroupStageBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var stages []client.Stage
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	target := "my-target"
	newTarget := "new-my-target"
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_disable_server_group_stage.s1"
	stage2 := "spinnaker_pipeline_disable_server_group_stage.s2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServerGroupOnlyDisabled,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineDisableServerGroupStageConfigBasic(pipeName, target, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target", target),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target", target),
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
				Config: testAccPipelineDisableServerGroupStageConfigBasic(pipeName, newTarget, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target", newTarget),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target", newTarget),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineDisableServerGroupStageConfigBasic(pipeName, target, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target", target),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineDisableServerGroupStageConfigBasic(pipeName, target, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{}, &stages),
				),
			},
		},
	})
}

func testAccPipelineDisableServerGroupStageConfigBasic(pipeName string, target string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_disable_server_group_stage" "s%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	cluster  = "test_cluster"
	target   = "%v"
}`, i, i, target)
	}

	return testAccPipelineConfigBasic("app", pipeName) + stages
}

func testAccCheckServerGroupOnlyDisabled(s *terraform.State) error {
	pipelineService := testAccProvider.Meta().(*Services).PipelineService
	for _, rs := range s.RootModule().Resources {
		if _, ok := stageTypes[rs.Type]; ok {
			_, err := pipelineService.GetPipelineByID(rs.Primary.Attributes[PipelineKey])
			if err != nil {
				return fmt.Errorf("Pipeline stage does not exist: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

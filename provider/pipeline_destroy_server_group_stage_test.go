package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func init() {
	stageTypes["spinnaker_pipeline_destroy_server_group_stage"] = client.DestroyServerGroupStageType
}

func TestAccPipelineDestroyServerGroupStageBasic(t *testing.T) {
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	target := "my-target"
	newTarget := "new-my-target"
	pipeline := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_destroy_server_group_stage.1"
	stage2 := "spinnaker_pipeline_destroy_server_group_stage.2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineStageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineDestroyServerGroupStageConfigBasic(pipeName, target, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target", target),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target", target),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				Config: testAccPipelineDestroyServerGroupStageConfigBasic(pipeName, newTarget, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target", newTarget),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target", newTarget),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				Config: testAccPipelineDestroyServerGroupStageConfigBasic(pipeName, target, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target", target),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
					}),
				),
			},
			{
				Config: testAccPipelineDestroyServerGroupStageConfigBasic(pipeName, target, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineStages(pipeline, []string{}),
				),
			},
		},
	})
}

func testAccPipelineDestroyServerGroupStageConfigBasic(pipeName string, target string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_destroy_server_group_stage" "%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	cluster  = "test_cluster"
	target   = "%v"
}`, i, i, target)
	}

	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "app"
	name        = "%s"
}`, pipeName) + stages
}

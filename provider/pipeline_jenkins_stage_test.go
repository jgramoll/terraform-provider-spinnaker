package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func init() {
	stageTypes["spinnaker_pipeline_jenkins_stage"] = client.JenkinsStageType
}

func TestAccPipelineJenkinsStageBasic(t *testing.T) {
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	master := "inst-ci"
	newMaster := master + "-new"
	pipeline := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_jenkins_stage.1"
	stage2 := "spinnaker_pipeline_jenkins_stage.2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineStageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineJenkinsStageConfigBasic(pipeName, master, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "master", master),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "master", master),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				ResourceName:      stage1,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      stage2,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPipelineJenkinsStageConfigBasic(pipeName, newMaster, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "master", newMaster),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "master", newMaster),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				Config: testAccPipelineJenkinsStageConfigBasic(pipeName, master, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "master", master),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
					}),
				),
			},
			{
				Config: testAccPipelineJenkinsStageConfigBasic(pipeName, master, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineStages(pipeline, []string{}),
				),
			},
		},
	})
}

func testAccPipelineJenkinsStageConfigBasic(pipeName string, master string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_jenkins_stage" "%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	master   = "%v"
	job      = "jenkins/job"
}`, i, i, master)
	}

	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "app"
	name        = "%s"
}`, pipeName) + stages
}

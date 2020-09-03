package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccPipelineExpectedArtifactsStageBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	reference := "my-ref"
	newReference := reference + "-new"
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_run_job_manifest_stage.stage"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineExpectedArtifactsStageBasicConfig(pipeName, reference, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "expected_artifact.0.default_artifact.0.custom_kind", "true"),
					resource.TestCheckResourceAttr(stage1, "expected_artifact.0.display_name", "my art 1"),
					resource.TestCheckResourceAttr(stage1, "expected_artifact.0.match_artifact.0.location", "loc"),
					resource.TestCheckResourceAttr(stage1, "expected_artifact.0.match_artifact.0.name", "name"),
					resource.TestCheckResourceAttr(stage1, "expected_artifact.0.match_artifact.0.reference", reference),
					resource.TestCheckResourceAttr(stage1, "expected_artifact.0.match_artifact.0.type", "s3/object"),
					resource.TestCheckResourceAttr(stage1, "expected_artifact.1.display_name", "my art 2"),
					resource.TestCheckResourceAttr(stage1, "expected_artifact.1.match_artifact.0.reference", reference),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
				),
			},
			{
				Config: testAccPipelineExpectedArtifactsStageBasicConfig(pipeName, newReference, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "expected_artifact.0.display_name", "my art 1"),
					resource.TestCheckResourceAttr(stage1, "expected_artifact.0.match_artifact.0.reference", newReference),
					resource.TestCheckResourceAttr(stage1, "expected_artifact.1.display_name", "my art 2"),
					resource.TestCheckResourceAttr(stage1, "expected_artifact.1.match_artifact.0.reference", newReference),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
				),
			},
			{
				Config: testAccPipelineExpectedArtifactsStageBasicConfig(pipeName, reference, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "expected_artifact.0.display_name", "my art 1"),
					resource.TestCheckResourceAttr(stage1, "expected_artifact.0.match_artifact.0.reference", reference),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
				),
			},
		},
	})
}

func testAccPipelineExpectedArtifactsStageBasicConfig(pipeName string, reference string, count int) string {
	artifacts := ""
	for i := 1; i <= count; i++ {
		artifacts += fmt.Sprintf(`
		expected_artifact {
			default_artifact {
				custom_kind = true
			}
			display_name = "my art %v"
			match_artifact {
				location = "loc"
				name = "name"
				reference = "%v"
				type = "s3/object"
			}
		}`, i, reference)
	}

	return testAccPipelineConfigBasic("app", pipeName) + fmt.Sprintf(`
resource "spinnaker_pipeline_run_job_manifest_stage" "stage" {
	pipeline    = spinnaker_pipeline.test.id
	name        = "Stage 1"
	account     = "acc"
	application = "my-app"

	cloud_provider = "my-cloud"
	source         = "text"

	manifest = <<EOT
first: 1
EOT

	%v
}`, artifacts)
}

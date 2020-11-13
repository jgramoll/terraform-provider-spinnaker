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
	stageTypes["spinnaker_pipeline_bake_manifest_stage"] = client.BakeManifestStageType
}

func TestAccPipelineBakeManifestStageBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var stages []client.Stage
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	namespace := "my-name"
	newNamespace := namespace + "-new"
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_bake_manifest_stage.s1"
	stage2 := "spinnaker_pipeline_bake_manifest_stage.s2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineBakeManifestStageConfigBasic(pipeName, namespace, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "namespace", namespace),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "namespace", namespace),
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
				Config: testAccPipelineBakeManifestStageConfigBasic(pipeName, newNamespace, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "namespace", newNamespace),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "namespace", newNamespace),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineBakeManifestStageConfigBasic(pipeName, namespace, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "namespace", namespace),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineBakeManifestStageConfigBasic(pipeName, namespace, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{}, &stages),
				),
			},
		},
	})
}

func TestAccPipelineBakeManifestStageKustomize(t *testing.T) {
	var pipelineRef client.Pipeline
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_bake_manifest_stage.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineBakeManifestStageConfigKustomize(pipeName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Build kustomize"),
					resource.TestCheckResourceAttr(stage1, "template_renderer", "KUSTOMIZE"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
				),
			},
		},
	})
}

func testAccPipelineBakeManifestStageConfigBasic(pipeName string, accountName string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_bake_manifest_stage" "s%v" {
	pipeline  = "${spinnaker_pipeline.test.id}"
	name      = "Stage %v"
	namespace = "%v"

	template_renderer = "HELM2"
}`, i, i, accountName)
	}

	return testAccPipelineConfigBasic("app", pipeName) + stages
}

func testAccPipelineBakeManifestStageConfigKustomize(pipeName string) string {
	stage := fmt.Sprintf(`
resource "spinnaker_pipeline_bake_manifest_stage" "test" {
	pipeline  = spinnaker_pipeline.test.id
	name      = "Build kustomize"

	template_renderer = "KUSTOMIZE"
	expected_artifact {
		default_artifact {
			custom_kind = true
		}
		display_name = "deploy-kustomize"
		match_artifact {
			artifact_account = "embedded-artifact"
			custom_kind = false
			type = "embedded/base64"
		}
		use_default_artifact = false
		use_prior_artifact = false
	}
	input_artifact {
		account = "test-account"
		artifact {
			artifact_account = "test-account"
			custom_kind = true
			metadata = {
				subPath = "checkout/path"
			}
			reference = "mygithub.repo"
			type = "git/repo"
			version = "github_auth"
		}
	}
	kustomize_file_path = "deploy/kustomize/nonprod/kustomization.yaml"
}`)

	return testAccPipelineConfigBasic("app", pipeName) + stage
}

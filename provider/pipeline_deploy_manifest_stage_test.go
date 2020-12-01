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
	stageTypes["spinnaker_pipeline_deploy_manifest_stage"] = client.DeployManifestStageType
}

func TestAccPipelineDeployManifestStageBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var stages []client.Stage
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	accountName := "my-account"
	newAccountName := accountName + "-new"
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_deploy_manifest_stage.s1"
	stage2 := "spinnaker_pipeline_deploy_manifest_stage.s2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineDeployManifestStageConfigBasic(pipeName, accountName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "account", accountName),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "account", accountName),
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
				Config: testAccPipelineDeployManifestStageConfigBasic(pipeName, newAccountName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "account", newAccountName),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "account", newAccountName),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineDeployManifestStageConfigBasic(pipeName, accountName, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "account", accountName),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineDeployManifestStageConfigBasic(pipeName, accountName, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{}, &stages),
				),
			},
		},
	})
}

func TestAccPipelineDeployManifestStageArtifact(t *testing.T) {
	var pipelineRef client.Pipeline
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_deploy_manifest_stage.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineDeployManifestStageConfigArtifact(pipeName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Deploy"),
					resource.TestCheckResourceAttr(stage1, "account", "test-account"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
				),
			},
		},
	})
}

func testAccPipelineDeployManifestStageConfigBasic(pipeName string, accountName string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_deploy_manifest_stage" "s%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	account  = "%v"

	cloud_provider            = "provider"
	source                    = "text"

	moniker {
		app = "app"
	}
	relationships {}
	traffic_management {
		options {}
	}

	manifests = [
		<<EOT
first: 1
EOT
,
<<EOT
second: 2
EOT
	]
}`, i, i, accountName)
	}

	return testAccPipelineConfigBasic("app", pipeName) + stages
}

func testAccPipelineDeployManifestStageConfigArtifact(pipeName string) string {
	stage := `
resource "spinnaker_pipeline_deploy_manifest_stage" "test" {
	pipeline = spinnaker_pipeline.test.id
	name     = "Deploy"
	account  = "test-account"

	cloud_provider = "kubernetes"
	source         = "artifact"

	manifest_artifact_id = "1234"
	skip_expression_evaluation = true

	moniker {
		app = "app"
	}
	relationships {}
	traffic_management {
		options {}
	}
}
`

	return testAccPipelineConfigBasic("app", pipeName) + stage
}

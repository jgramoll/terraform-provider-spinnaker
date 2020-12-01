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
	stageTypes["spinnaker_pipeline_evaluate_variables_stage"] = client.EvaluateVariablesStageType
}

func TestAccPipelineEvaluateVariablesStageBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var stages []client.Stage
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	variables := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	newVariables := map[string]string{
		"foo": "bar",
		"baz": "quux",
	}
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_evaluate_variables_stage.s1"
	stage2 := "spinnaker_pipeline_evaluate_variables_stage.s2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineEvaluateVariablesStageConfigBasic(pipeName, variables, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "variables.foo", "bar"),
					resource.TestCheckResourceAttr(stage1, "variables.baz", "qux"),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "variables.foo", "bar"),
					resource.TestCheckResourceAttr(stage2, "variables.baz", "qux"),
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
				Config: testAccPipelineEvaluateVariablesStageConfigBasic(pipeName, newVariables, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "variables.foo", "bar"),
					resource.TestCheckResourceAttr(stage1, "variables.baz", "quux"),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "variables.foo", "bar"),
					resource.TestCheckResourceAttr(stage2, "variables.baz", "quux"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineEvaluateVariablesStageConfigBasic(pipeName, variables, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "variables.foo", "bar"),
					resource.TestCheckResourceAttr(stage1, "variables.baz", "qux"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineEvaluateVariablesStageConfigBasic(pipeName, variables, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{}, &stages),
				),
			},
		},
	})
}

func testAccPipelineEvaluateVariablesStageConfigBasic(pipeline string, variables map[string]string, count int) string {
	variablesString := ""
	for k, v := range variables {
		variablesString += fmt.Sprintf(`
	%s = "%s"
`, k, v)
	}

	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_evaluate_variables_stage" "s%v" {
	pipeline 	  = "${spinnaker_pipeline.test.id}"
	name     	  = "Stage %v"

	variables = {%v}
}`, i, i, variablesString)
	}

	return testAccPipelineConfigBasic("app", pipeline) + stages
}

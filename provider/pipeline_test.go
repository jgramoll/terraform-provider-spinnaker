package provider

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccKubernetesSecret_basic(t *testing.T) {
	var pipeline client.Pipeline
	name := fmt.Sprintf("tf-acc-test-%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	fmt.Println("basic test", conf)
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: "spinnaker_pipeline.test",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckPipelineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists("spinnaker_pipeline.test", &pipeline),
					resource.TestCheckResourceAttr("spinnaker_pipeline.test", "application", "asdf"),
					testAccCheckPipeline(&pipeline, map[string]string{"TestAnnotationOne": "one", "TestAnnotationTwo": "two"}),
				),
			},
		},
	})
}

func testAccPipelineConfigBasic(name string) string {
	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "#app"
	name        = "%s"
}
`, name)
}

func testAccCheckPipelineExists(n string, obj *client.Pipeline) resource.TestCheckFunc {
	fmt.Println("testAccCheckPipelineExists")
	return func(s *terraform.State) error {
		fmt.Println("testAccCheckPipelineExistsFrd")
		// fmt.Println("testAccCheckPipelineExists", s)
		fmt.Println("s", s)
		fmt.Println("s", s.RootModule())
		fmt.Println("s", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			fmt.Println("boogy", n)
			return fmt.Errorf("Not found: %s", n)
		}
		fmt.Println(rs)

		log.Println(rs)

		// conn := testAccProvider.Meta().(*kubernetes.Clientset)

		// namespace, name, err := idParts(rs.Primary.ID)
		// if err != nil {
		// 	return err
		// }

		// out, err := conn.CoreV1().Secrets(namespace).Get(name, meta_v1.GetOptions{})
		// if err != nil {
		// 	return err
		// }

		// *obj = *out
		return nil
	}
}

func testAccCheckPipeline(p *client.Pipeline, expected map[string]string) resource.TestCheckFunc {
	fmt.Println("testAccCheckPipeline")
	return func(s *terraform.State) error {
		log.Println("testAccCheckPipelineFrd")
		// 	if len(expected) == 0 && len(om.Annotations) == 0 {
		// 		return nil
		// 	}

		// 	// Remove any internal k8s annotations unless we expect them
		// 	annotations := om.Annotations
		// 	for key, _ := range annotations {
		// 		_, isExpected := expected[key]
		// 		if isInternalKey(key) && !isExpected {
		// 			delete(annotations, key)
		// 		}
		// 	}

		// 	if !reflect.DeepEqual(annotations, expected) {
		// 		return fmt.Errorf("%s annotations don't match.\nExpected: %q\nGiven: %q",
		// 			om.Name, expected, om.Annotations)
		// 	}
		return nil
	}
}

func testAccCheckPipelineDestroy(s *terraform.State) error {
	// conn := testAccProvider.Meta().(*client.Client)

	log.Println("testAccCheckPipelineDestroy")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kubernetes_namespace" {
			continue
		}

		// resp, err := conn.CoreV1().Namespaces().Get(rs.Primary.ID, meta_v1.GetOptions{})
		// if err == nil {
		// 	if resp.Name == rs.Primary.ID {
		// 		return fmt.Errorf("Namespace still exists: %s", rs.Primary.ID)
		// 	}
		// }
	}

	return nil
}

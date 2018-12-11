package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

var p *schema.Resource
var c *client.Client

func init() {
	p = pipelineResource()
	c = client.NewClient((client.Config{
		Address:  "",
		CertPath: "",
		KeyPath:  ""}))
}

func TestPipelineRead(t *testing.T) {
	d := &schema.ResourceData{}
	err := p.Read(d, c)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(d)
}

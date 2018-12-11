package provider

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func pipelineResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"application": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	// name := d.Get("name").(string)
	// d.SetId(name)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)

	pipeline, err := c.GetPipeline("TODO.appName", d.Id())
	if err != nil {
		log.Printf("[WARN] No Server found: %s", d.Id())
		d.SetId("")
		return nil
	}

	log.Printf("[INFO] got pipeline %s", pipeline.Id)
	d.Set("name", pipeline.Name)
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	// Enable partial state mode
	d.Partial(true)

	// if d.HasChange("address") {
	//   // Try updating the address
	//   if err := updateAddress(d, m); err != nil {
	//           return err
	//   }

	//   d.SetPartial("address")
	// }

	d.Partial(false)

	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

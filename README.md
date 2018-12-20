# terraform-provider-spinnaker
Terraform Provider to manage spinnaker pipelines

## Build and install ##

### Dependencies ###

You should have a working Go environment setup.  If not check out the Go [getting started](http://golang.org/doc/install) guide.

[Glide](https://github.com/Masterminds/glide) is used for dependency management.  To install all dependencies run `glide i`.

### Install ###

You will need to install the binary as a [terraform third party plugin](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).  Terraform will then pick up the binary from the local filesystem when you run `terraform init`.

```sh
ln -s ~/go/bin/terraform-provider-spinnaker ~/.terraform.d/plugins/$(uname | tr '[:upper:]' '[:lower:]')_amd64/terraform-provider-spinnaker_v$(date +%Y.%m.%d)
```

## Usage ##

```terraform
resource "spinnaker_pipeline" "edge" {
	application = "career"
	name        = "My New Pipeline"
  index       = 4
}

resource "spinnaker_pipeline_trigger" "jenkins" {
	pipeline = "${spinnaker_pipeline.edge.id}"
	job = "Bridge Career/job/Bridge_nav/job/Bridge_nav_postmerge"
	master = "inst-ci"
	property_file = "build.properties.test"
	type = "jenkins"
}

resource "spinnaker_pipeline_notification" "edge" {
	pipeline = "${spinnaker_pipeline.edge.id}"
	address = "bridge-career-deploys"
	level = "pipeline"
	message = {
		complete = "edge is done"
		failed = "edge is failed"
		starting = "edge is starting"
	}
	type = "slack"
	when = {
		complete = true
		starting = false
		failed = true
	}
}

resource "spinnaker_pipeline_bake_stage" "bake" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage Bake"
}
```

## TODO

1. Notifications
1. Parameters
1. Stages
1. Import

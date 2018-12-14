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
resource "spinnaker_pipeline" "test" {
	application = "career"
	name        = "My New Pipeline"
  index       = 4
}

resource "spinnaker_pipeline_trigger" "jenkins" {
	pipeline = "${spinnaker_pipeline.test.id}"
	job = "Bridge Career/job/Bridge_nav/job/Bridge_nav_postmerge"
	master = "inst-ci"
	property_file = "build.properties.test"
	type = "jenkins"
}
```

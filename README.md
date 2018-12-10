# terraform-provider-spinnaker
Terraform Provider to manage spinnaker pipelines

## Build and install ##

### Dependencies ###

You should have a working Go environment setup.  If not check out the Go [getting started](http://golang.org/doc/install) guide.

This use [dep](https://github.com/golang/dep) for dependency management.  To fetch all dependencies run `dep ensure`.

### Install ###

You will need to install the binary as a [terraform third party plugin](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).  Terraform will then pick up the binary from the local filesystem when you run `terraform init`.

```
ln -s ~/go/bin/terraform-provider-spinnaker ~/.terraform.d/plugins/$(uname | tr '[:upper:]' '[:lower:]')_amd64/terraform-provider-spinnaker_v$(date +%Y.%m.%d)
```

## Usage ##

TODO

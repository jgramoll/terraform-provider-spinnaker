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
  parameter {
    name = "My Parameter"
  }

  parameter {
    name        = "Detailed Parameter"
    description = "Setting options"
    default     = "Default value"
    label       = "Trigger label"

    option {
      value = 1
    }
    option {
      value = "two"
    }
  }
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

resource "spinnaker_pipeline_jenkins_stage" "bake" {
  pipeline = "${spinnaker_pipeline.test.id}"
  name     = "Stage Jenkins"

  requisite_stage_ref_ids = ["${spinnaker_pipeline_bake_stage.id}"]

  notification {
    address = "#my-slack-channel"
    message = {
      failed = "Jenkins Stage failed"
    }
    type = "slack"
    when = {
      failed = true
    }
  }
}

resource "spinnaker_pipeline_deploy_stage" "deploy" {
  pipeline = "${spinnaker_pipeline.test.id}"
  name     = "Stage Deploy"
  restricted_execution_window {
    days = [1,3]
    jitter {
      enabled = true
      max_delay = 5
    }
    whitelist {
      end_hour = 1
      end_min = 2
    }
    whitelist {
      end_hour = 3
      end_min = 4
    }
  }
  cluster {
    account = "my-account"
    application = "app"
    availability_zones {
      us_east_1 = [
        "us-east-1a",
        "us-east-1b",
        "us-east-1c"
      ]
    }
    capacity {
      desired = 2
    }
    cloud_provider = "aws"
    health_check_type = "ELB"
    instance_type = "t2.micro"
    key_pair = "my_key_pair"
    provider = "aws"
    strategy = "redblack"
    subnet_type = "my_subnet"
  }
  cluster {
    account = "my-account"
    application = "app"
    availability_zones {
      us_east_2 = [
        "us-east-2a",
        "us-east-2b",
        "us-east-2c"
      ]
    }
    cloud_provider = "aws"
    health_check_type = "ELB"
    instance_type = "t2.micro"
    key_pair = "my_key_pair"
    provider = "aws"
    strategy = "redblack"
    subnet_type = "my_subnet"
  }
}

resource "spinnaker_pipeline_rollback_cluster_stage" "deploy" {
  pipeline = "${spinnaker_pipeline.test.id}"
  name     = "Rollback Cluster"

  cloud_provider      = "aws"
  cloud_provider_type = "aws"
  cluster             = "my-cluster"
  credentials         = "my-creds"
  moniker {
    app     = "my-app"
    cluster = "my-cluster"
    detail  = "api"
    stack   = "edge"
  }
  regions = [
    "us-east-2"
  ]
}

resource "spinnaker_pipeline_destroy_server_group_stage" "deploy" {
  pipeline = "${spinnaker_pipeline.test.id}"
  name     = "Destroy Server Group"

  cloud_provider      = "aws"
  cloud_provider_type = "aws"
  cluster             = "my-cluster"
  credentials         = "my-creds"
  moniker {
    app     = "my-app"
    cluster = "my-cluster"
    detail  = "api"
    stack   = "edge"
  }
  regions = [
    "us-east-2"
  ]
  target = "oldest_asg_dynamic"
}

resource "spinnaker_pipeline_pipeline_stage" "main" {
  pipeline = "${spinnaker_pipeline.test.id}"
  name     = "Start prod deploy"

  application      = "my-other-app"
  target_pipeline  = "${spinnaker_pipeline.prod.id}"
  pipelineParameters {
    version = "my-message"
  }
}

```

## TODO

1. Parameters
1. More Stages
1. Stage Notifications
1. More Triggers (refactor to make polymorphic)
1. Import

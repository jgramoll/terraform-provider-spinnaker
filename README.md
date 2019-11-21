# terraform-provider-spinnaker
Terraform Provider to manage spinnaker pipelines

## Install ##

You will need to install the binary as a [terraform third party plugin](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).  Terraform will then pick up the binary from the local filesystem when you run `terraform init`.

```sh
curl -s https://raw.githubusercontent.com/jgramoll/terraform-provider-spinnaker/master/install.sh | bash
```

## Usage ##

### Credentials ###

```sh
export SPINNAKER_ADDRESS=https://your.spinnaker.server
export SPINNAKER_CERT=/path/to/spinnaker/cert
export SPINNAKER_KEY=/path/to/spinnaker/key
export SPINNAKER_EMAIL=your@email.org
```

```terraform
provider "spinnaker" {
  address = "${var.spinnaker_address}"

  cert_path  = "${var.cert_path}"
  key_path   = "${var.key_path}"
  user_email = "${var.user_email}"
}

resource "spinnaker_application" "application" {
  name             = "myApplication"
  email            = "testo@test.com"
  repo_type        = "github"
  repo_project_key = "http://github.com/user/my-project"
  repo_slug        = "my-project"

  cloud_providers = [
    "aws",
  ]

  provider_settings = {
    aws = {
      use_ami_block_device_mappings = true
    }
  }

  instance_port = 8080
}

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

resource "spinnaker_pipeline_docker_trigger" "docker" {
  pipeline = "${spinnaker_pipeline.edge.id}"

  account = "my-docker-hub"
  organization = "my-org"
  repository = "test"
}

resource "spinnaker_pipeline_jenkins_trigger" "jenkins" {
  pipeline = "${spinnaker_pipeline.edge.id}"
  job = "Bridge Career/job/Bridge_nav/job/Bridge_nav_postmerge"
  master = "inst-ci"
  property_file = "build.properties.test"
}

resource "spinnaker_pipeline_pipeline_trigger" "pipeline" {
  pipeline = "${spinnaker_pipeline.edge.id}"

  triggering_application = "app"
  triggering_pipeline = "my-other-pipeline"
  status = ["successful"]
}

resource "spinnaker_pipeline_notification" "edge" {
  pipeline = "${spinnaker_pipeline.edge.id}"
  address = "bridge-career-deploys"
  message {
    complete = "edge is done"
    failed = "edge is failed"
    starting = "edge is starting"
  }
  type = "slack"
  when {
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

  requisite_stage_ref_ids = ["${spinnaker_pipeline_bake_stage.bake.id}"]

  notification {
    address = "#my-slack-channel"
    message {
      failed = "Jenkins Stage failed"
    }
    type = "slack"
    when {
      failed = true
    }
  }
}

resource "spinnaker_pipeline_manual_judgment_stage" "main" {
  pipeline     = "${spinnaker_pipeline.test.id}"
  name         = "Judgment"
  instructions = "Manual Judgment Instructions"

  judgment_inputs = [
    "commit",
    "rollback",
  ]
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

resource "spinnaker_pipeline_find_image_from_tags_stage" "main" {
  pipeline     = "${spinnaker_pipeline.test.id}"
  name         = "Find image"
  package_name = "my-package"

  cloud_provider      = "aws"
  cloud_provider_type = "aws"
}

resource "spinnaker_pipeline_resize_server_group_stage" "deploy" {
  pipeline = "${spinnaker_pipeline.test.id}"
  name     = "Resize Server Group"

  action              = "scale_exact"
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
  resize_type = "exact"
  target = "oldest_asg_dynamic"
}

resource "spinnaker_pipeline_pipeline_stage" "main" {
  pipeline = "${spinnaker_pipeline.test.id}"
  name     = "Start prod deploy"

  application      = "my-other-app"
  target_pipeline  = "${spinnaker_pipeline.prod.id}"
  pipeline_parameters {
    version = "my-message"
  }
}

resource "spinnaker_pipeline_parameter" "version" {
  pipeline = "${spinnaker_pipeline.test.id}"
  name     = "version"
}

resource "spinnaker_pipeline_delete_manifest_stage" "main" {
  pipeline = "${spinnaker_pipeline.test.id}"
  name     = "Delete manifest"
  account  = "account"
  app      = "app"

  cloud_provider = "kubernetes"
  location       = "location"
  manifest_name  = "manifest name"
  mode           = "static"
}

resource "spinnaker_pipeline_deploy_manifest_stage" "main" {
  pipeline = "${spinnaker_pipeline.test.id}"
  name     = "Deploy Manifest"
  account  = "account"

  cloud_provider            = "kubernetes"
  source                    = "text"
  manifest_artifact_account = "docker-registry"

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
}

resource "spinnaker_pipeline_webhook_stage" "main" {
  pipeline = "${spinnaker_pipeline.test.id}"
  name     = "Webhook"
  url      = "http://my-webhook.io/"
}

resource "spinnaker_canary_config" "test" {
  name         = "%s"
  applications = ["app"]
  metric {
    groups = ["Group 1"]
    name = "my metric"
    query {
      metric_name = "avg:aws.ec2.cpucredit_balance"
      service_type = "datadog"
      type = "datadog"
    }
  }
  classifier {
    group_weights = {
      "Group 1" = 100
    }
  }
  judge {
    name = "NetflixACAJudge-v1.0"
  }
}

resource "spinnaker_pipeline_canary_analysis_stage" "test" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Canary Analysis"

	analysis_type  = "realTimeAutomatic"

	canary_config {
		canary_config_id  = "${spinnaker_canary_config.test.id}"
		lifetime_duration = "PT0H5M"

		metrics_account_name = "metrics-account"
		storage_account_name = "storage-account"

		score_thresholds {
			marginal = "1"
			pass     = "2"
		}
	}
}

```

## Local Dev ##

### Depenedencies ###

You should have a working Go environment setup.  If not check out the Go [getting started](http://golang.org/doc/install) guide.

[Go modules](https://github.com/golang/go/wiki/Modules) are used for dependency management.  To install all dependencies run the following:

```sh
export GO111MODULE=on
go mod vendor
```

### Link ###

```sh
go clean
go build
rm ~/.terraform.d/plugins/$(uname | tr '[:upper:]' '[:lower:]')_amd64/terraform-provider-spinnaker_v2.0.0
ln  ./terraform-provider-spinnaker ~/.terraform.d/plugins/$(uname | tr '[:upper:]' '[:lower:]')_amd64/terraform-provider-spinnaker_v2.0.0
```

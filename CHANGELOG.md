# terraform-provider-spinnaker Change Log

## 1.2.0

- `564f3ce` Deploy manifest (#20)
- `606710d` fix state if pipeline is missing to not be fatal (#19)

## 1.1.1

- `e6c94d5` fix stage read to not error if stage was deleted remotely (#17)

## 1.1.0

- `c251969` allow any AZ not just us-east-1, us-east-2
- `5e0fd8e` fix notification when null pointer exception
- `2d2a230` fix terraform plugins dir to use kernel (#14)
- `a762635` just map all the known regions
- `a7ef27f` nit remove whitespace
- `e3a21f8` update goreleaser to prerelease:auto

## 1.0.0

- `46ffe9c` Adds pipeline parameter resource (#4)
- `15469ed` Importer (#6)
- `14360b1` Initial commit
- `6837558` Merge pull request #3 from jgramoll/stageNotifications
- `67f796c` add ability to add notifications to a specific pipeline stage
- `8c1d17b` add application CRUD to client
- `e99452a` add bake package (#12)
- `6234a6b` add deploy stage
- `3989a3a` add destroy server group stage (#2)
- `1db4904` add goreleaser config (#13)
- `1c233bc` add initial acceptance tests
- `6328ede` add intial client to send http requests
- `eee5103` add jenkins stage
- `9994736` add more validation to notification resource schema
- `3273c48` add notifications
- `24f3504` add parameters and resize resources (#7)
- `38d9105` add pipeline trigger stage
- `b88c70b` add rollback cluster stage
- `f148dff` add target_healthy_deploy_percentage to rollback_cluster_stage (#9)
- `525cffe` add tests for new pipeline stage/notification methods
- `af4890a` add trigger resource
- `735046d` allow provider to update pipeline
- `a67416c` allow setting deploy.max_remaining_asgs (#11)
- `8f8e552` clean up some code smells / duplication
- `e30365f` debug acceptance test
- `145faa2` finish stage refactor and get acceptance tests working again
- `ec4112f` fix jenkins params to allow bools (#10)
- `1cdcb3c` fix main spacing
- `601ba8b` fix pipeline param option
- `02429e2` fix pipeline params
- `436c623` fix pipeline stage resource types
- `6d631b8` fix provider/pipeline client call
- `d1bc581` fix readme pipeline stage params
- `f3ed5e4` get create/read/delete working for pipeline resource
- `72c9dcf` hook up client.GetPipeline
- `039584a` initial project with empty provider
- `a7941ae` massive refactor of client
- `2ccff76` move stage crud and notification crud to client.Pipeline
- `1c4d221` refactor notification when to not need to specify 'pipeline'
- `2db0fa3` refactor pipeline trigger to seperate resource code
- `97c327d` seperate spinnaker provider and client
- `7d86129` standardize base stage to have notifications, stage_enabled
- `9ad1f1f` update default values to match spinnaker ui
- `94121ef` update docs about glide as dependency manager
- `a47ac10` update interfaces
- `0020c93` use gofmt to standardize formating

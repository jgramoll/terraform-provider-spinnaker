# terraform-provider-spinnaker

## 2.3.0

- [341ee64] Patch manifest (#72)
- [ec0d831] add bake manifest stage (#71)
- [cbb5078] Stage expected artifacts (#70)
- [6d185f7] Initializing Variables (#67)

## 2.2.2

- `2c8d5f` persist StageTimeoutMS (#66)

## 2.2.1

- `2671fd` fix goreleaser archive->archives

## 2.2.0

- `c67b2c6` Webhook trigger (#64)
- `3ac5f60` update application.accounts to be computed for better diffs (#62)
- `6fb2973` make sure deploy_manifest_stage can set credentials (#61)

## 2.1.0

- `ab4d008` Enable disable manifest (#59)
- `bfcb538` add undo rollout and enable server group stages (#58)

## 2.0.0

Breaking change:
`spinnaker_pipeline_deploy_manifest_stage` now requires `moniker` field.
Spinnaker was causing null pointer exceptions if it wasn't provided

- `70aa566` make sure deploy manifest moniker is required (#56)
- `3ed0849` add notifications to manual judgement (#55)

## 1.8.0

- `9624ffe` Evalaute variables (#52)
- `06e005c` add spinnaker_pipeline_find_artifacts_from_resource_stage (#51)
- `80179b9` fix application import crashing if no permissions defined (#50)

## 1.7.0

- `ffc3a41` add run_job_manifest to client (#48)
- `e3db9c8` Precondition (#47)
- `c978d77` add permissions to applications (#46)
- `ee4c984` refactor base stages (#44)
- `4b5f0be` Andrein traffic management (#43)
- `8dfb546` override namespace (#41)
- `7675b11` scale manifest (#42)
- `446e472` add pipeline ui locked (#38)
- `2960a84` Add registry and tag to docker trigger (#37)

## 1.6.0

- `331273b` bump terraform dependency to 0.12.16 (#34)
- `83a089c` Dockertrigger (#33)

## 1.5.0

- `48c4729` add canary config to client (#30)

## 1.4.0

- `a8a20e4` add pipeline date source (#29)

## 1.3.1

- `60bcb31` allow for pipeline index to be computed (#27)
- `4b067ab` add more resources to README (#26)

## 1.3.0

- `8d5a409` merge from fanatics (#24)
- `1ffae34` update terraform dependency to 0.12.6 (#8)

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

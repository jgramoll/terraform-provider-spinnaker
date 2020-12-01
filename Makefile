VERSION := v4.0.0

tag:
	@git tag -a ${VERSION} -m ${VERSION}
	@git push origin ${VERSION}

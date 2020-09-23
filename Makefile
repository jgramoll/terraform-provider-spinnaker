version=v3.1.0

tag:
	@git tag -a "${version}" -m "${version}"
	@git push origin "${version}"	

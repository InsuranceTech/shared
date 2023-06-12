#!make
SERVICE_NAME = shared

tag:
	@echo "Write the version (ex v1.0.47) : "; \
    read VERSION;
	git tag $(VERSION)
	git push origin $(VERSION)
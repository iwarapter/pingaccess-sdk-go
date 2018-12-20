# Makefile

# SERVICES = applications rules virtualhosts identityMappings siteAuthenticators sites
SERVICES := $(shell find models -name '*.json')

test:
	@rm report.json coverage.out
	@go test ./... -v -coverprofile=coverage.out -json > report.json

install:
	@go install ./...

generate:
	@find service/ ! -name '*_test.go' ! -name '*.json' -type f -exec rm -f {} +
	@/Users/iwarapter/golang-learning/terraform-provider-pingaccess/bin/pingaccess-sdk-go-gen-cli $(SERVICES)
	@go fmt ./...

sonar:
	@sonar-scanner \
		-Dsonar.projectKey=github.com.iwarapter.pingaccess-sdk-go \
		-Dsonar.sources=. \
		-Dsonar.host.url=http://localhost:9001 \
		-Dsonar.login=28d86a90f2e4ae9563b4501cbc99de7522219c88 \
		-Dsonar.go.coverage.reportPaths=coverage.out \
		-Dsonar.go.tests.reportPaths=report.json \
		-Dsonar.exclusions=vendor
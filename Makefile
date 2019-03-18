# Makefile

# SERVICES = applications rules virtualhosts identityMappings siteAuthenticators sites
SERVICES := $(shell find models -name '*.json')

test:
	@rm -f report.json coverage.out
	@go test ./... -v -coverprofile=coverage.out -json > report.json

install:
	@go install ./...

generate:
	@find pingaccess/ ! -name '*_test.go' ! -name 'pingaccess.go' ! -name 'pingaccess.lic' ! -name '*.json' -type f -exec rm -f {} +
	@$$GOPATH/bin/pingaccess-sdk-go-gen-cli $(SERVICES)
	@go fmt ./...

sonar:
	@sonar-scanner \
		-Dsonar.projectKey=github.com.iwarapter.pingaccess-sdk-go \
		-Dsonar.sources=. \
		-Dsonar.host.url=http://localhost:9001 \
		-Dsonar.login=28d86a90f2e4ae9563b4501cbc99de7522219c88 \
		-Dsonar.go.coverage.reportPaths=coverage.out \
		-Dsonar.go.tests.reportPaths=report.json \
		-Dsonar.exclusions=vendor \
		-Dsonar.tests="." \
		-Dsonar.test.inclusions="**/*_test.go"
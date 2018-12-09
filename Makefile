# Makefile

SERVICES = applications rules virtualhosts identityMappings siteAuthenticators sites

test:
	$(foreach var,$(SERVICES),go test ./service/$(var) -v -coverprofile=coverage-service-$(var).out && go tool cover -func=coverage-service-$(var).out;)

integration:
	$(foreach var,$(SERVICES),go test ./service/$(var) -v --tags integration -coverprofile=coverage-service-integration-$(var).out && go tool cover -func=coverage-service-integration-$(var).out;)
	
install:
	@go install ./...

generate:
	@find service/ ! -name '*_test.go' ! -name '*.json' -type f -exec rm -f {} +
	$(foreach var,$(SERVICES),/Users/iwarapter/golang-learning/terraform-provider-pingaccess/bin/pingaccess-sdk-go-gen-cli models/$(var).json;)
	@go fmt ./...

sonar:
	@sonar-scanner \
		-Dsonar.projectKey=github.com.iwarapter.pingaccess-sdk-go \
		-Dsonar.sources=. \
		-Dsonar.host.url=http://localhost:9001 \
		-Dsonar.login=28d86a90f2e4ae9563b4501cbc99de7522219c88 \
		-Dsonar.go.coverage.reportPaths=coverage.txt \
		-Dsonar.go.tests.reportPaths=report.json \
		-Dsonar.exclusions=vendor
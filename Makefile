# Makefile

SERVICES = applications rules virtualhosts identityMappings siteAuthenticators

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
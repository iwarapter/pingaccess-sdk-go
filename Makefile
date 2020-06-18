# Makefile

# SERVICES = applications rules virtualhosts identityMappings siteAuthenticators sites
SERVICES := $(shell find models -name '*.json')

pa-init:
	@docker run --rm -d --hostname pingaccess --name pingaccess -e PING_IDENTITY_DEVOPS_KEY=$(PING_IDENTITY_DEVOPS_KEY) -e PING_IDENTITY_DEVOPS_USER=$(PING_IDENTITY_DEVOPS_USER) -e PING_IDENTITY_ACCEPT_EULA=YES --publish 9000:9000 pingidentity/pingaccess:6.0.1-edge

test:
	@rm -f report.json coverage.out
	@go test ./... -v -coverprofile=coverage.out && go tool cover -func=coverage.out

test-and-report:
	@rm -f report.json coverage.out
	@go test ./... -v -coverprofile=coverage.out -json > report.json && go tool cover -func=coverage.out

checks:
	@go fmt ./...
	@staticcheck ./...
	@gosec ./...
	@goimports -w services pingaccess

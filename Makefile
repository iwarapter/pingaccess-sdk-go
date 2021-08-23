# Makefile

pa-init:
	@docker run --rm -d --hostname pingaccess --name pingaccess -e PING_IDENTITY_DEVOPS_KEY=$(PING_IDENTITY_DEVOPS_KEY) -e PING_IDENTITY_DEVOPS_USER=$(PING_IDENTITY_DEVOPS_USER) -e PING_IDENTITY_ACCEPT_EULA=YES --publish 9000:9000 pingidentity/pingaccess:6.0.2-edge

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

docs:
	docker run \
        --rm \
        -e "GOPATH=/tmp/go" \
        -p 127.0.0.1:6060:6060 \
        -v `pwd`:/tmp/go/src/github.com/iwarapter/pingaccess-sdk \
        --name godoc \
        golang \
        bash -c "go get golang.org/x/tools/cmd/godoc && echo http://localhost:6060/pkg/ && /tmp/go/bin/godoc -http=:6060"

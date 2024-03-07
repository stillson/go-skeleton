

build:
	go build ./...

test:
	go test

cover:
	go test -coverprofile coverage.out

report: cover
	go tool cover -html=coverage.out -o cover.html

check-format:
	test -z $$(go fmt ./...)

check: check-format static-check
	go vet ./...

#setup: install-go init-go install-lint
setup: install-lint

install-lint:
	sudo curl -sSfL \
	https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh\
	| sh -s -- -b $$(go env GOPATH)/bin latest

static-check:
	golangci-lint run

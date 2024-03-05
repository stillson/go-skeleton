

build:
	go build

test:
	go test

cover:
	go test -coverprofile coverage.out

report:
	go tool cover -html=coverage.out -o cover.html

check-format:
	test -z $$(go fmt ./...)

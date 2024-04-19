.PHONY: build
build:
	go build

.PHONY: install
install:
	go install github.com/simonmittag/puri/cmd/puri

.PHONY: test
test:
	go clean -testcache && go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: all
all: lint build test install


# Set up tools
.PHONY: install
install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.2

.PHONY: lint
lint:
	go env -w GOFLAGS=-buildvcs=false
	golangci-lint run .

.PHONY: run
run:
	go run philosophers.go

.PHONY: build
build:
	go mod tidy
	go mod vendor
	go build -v -o build/philosophers

.PHONY: clean
clean:
	go clean -v
	rm -rf build
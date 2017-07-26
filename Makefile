COMMIT = $$(git describe --always)

BUILD_FLAGS = -ldflags "-X main.Commit=\"$(COMMIT)\""

test:
	@echo "Run tests"
	@go test -v ./...

build:
	@echo "Build isbn-gen in ./bin"
	@go build $(BUILD_FLAGS) -o ./bin/isbn-gen

install:
	@echo "Install isbn-gen in $(GOPATH)/bin"
	@go install $(BUILD_FLAGS)

clean:
	@echo "Clean up ./bin directory"
	@rm -rf bin/*

.PHONY: test build install

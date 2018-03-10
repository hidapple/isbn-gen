COMMIT = $$(git describe --always)

BUILD_FLAGS = -ldflags "-X main.Commit=\"$(COMMIT)\""

test:
	@echo "===> Running tests..."
	@go test -v ./...

build:
	@echo "===> Building isbn-gen in ./bin directory..."
	@go build $(BUILD_FLAGS) -o ./bin/isbn-gen

cross-build:
	@echo "===> Building for cross platform..."
	@rm -rf ./dist
	@gox -os="linux darwin windows" -arch="386 amd64" -output "dist/isbn-gen_{{.OS}}_{{.Arch}}"

install:
	@echo "===> Installing isbn-gen in $(GOPATH)/bin directory..."
	@go install $(BUILD_FLAGS)

clean:
	@echo "===> Cleaning up ./bin directory..."
	@rm -rf bin/*

.PHONY: test build cross-build install clean

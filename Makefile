GOFMT?=gofmt "-s"
BINARY_NAME=lightgate
VERSION=$(shell git describe --tags --abbrev=0)
BUILD_RELEASE=go build -ldflags="-s -w -X main.version=$(VERSION)" -o
BUILD_DEV=go build -o
BUILD_DIR=./build

.PHONY: fmt
fmt:
	@echo "Formatting & simplifying Go code..."
	@$(GOFMT) -w .
	@echo "Done!"

.PHONY: check-fmt
check-fmt:
	@echo "Checking code formatting..."
	@output=$$($(GOFMT) -l .); \
	if [ -n "$$output" ]; then \
		echo "Files not formatted:"; \
		echo "$$output"; \
		exit 1; \
	else \
		echo "All files are properly formatted"; \
	fi

build-dev:
	$(BUILD_DEV) $(BINARY_NAME) lightgate.go

build-win:
	GOOS=windows GOARCH=amd64 $(BUILD_RELEASE) $(BUILD_DIR)/$(BINARY_NAME)$(VERSION)-win.exe lightgate.go

build-linux:
	GOOS=linux GOARCH=amd64 $(BUILD_RELEASE) $(BUILD_DIR)/$(BINARY_NAME)$(VERSION)-linux lightgate.go

build-macos:
	GOOS=darwin GOARCH=amd64 $(BUILD_RELEASE) $(BUILD_DIR)/$(BINARY_NAME)$(VERSION)-macos lightgate.go

build-all: build-win build-linux build-macos
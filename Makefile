GOFMT ?= gofmt "-s"
BUILD = go build -ldflags="-s -w" -trimpath -o

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

build:
	$(BUILD) lightgate lightgate.go

build-win:
	$(BUILD) lightgate.exe lightgate.go
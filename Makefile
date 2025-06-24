
PLUGIN_NAME := packer-plugin-syft
VERSION := 1.0.0

.PHONY: build
build:
	go build -o $(PLUGIN_NAME) .

.PHONY: install
install: build
	mkdir -p ~/.packer.d/plugins
	cp $(PLUGIN_NAME) ~/.packer.d/plugins/

.PHONY: clean
clean:
	rm -f $(PLUGIN_NAME)

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: plugin-check

plugin-check:
	@echo "Running plugin checks..."
	# Add the actual commands here that you want to run for plugin checking.
	# For example, if you have a custom Go program to run checks:
	# go run ./internal/plugins/checker/main.go
	# Or a shell script:
	# ./scripts/check-plugins.sh
	@echo "Plugin checks complete."

# ... other Makefile targets (build, test, clean, etc.) ...

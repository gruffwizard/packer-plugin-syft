
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

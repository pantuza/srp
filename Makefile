#
# Makefile to rule project build and execution
#


# Protocol Buffers binary
PROTOC := $(shell which protoc)

# Protocol Buffers plugin for Go Lang
PROTOC_GO := $(shell which protoc-gen-go)


# Golang binary path
GO := $(shell which go)


default: help


help:
	@echo "Service Ready Protocol target rules:"
	@echo
	@echo "build           Builds the project. Generates binary"
	@echo "run             Runs the main code using standart input and output"
	@echo "test            Runs the service ready protocol tests"
	@echo "clean           Cleans the entire project"
	@echo "proto           Builds protocol buffers and gRPC classes"
	@echo "fmt             Formats code using go fmt"
	@echo "help            Prints this help message"


build: main.go
	$(GO) build $^


run: main.go
	$(GO) run $^


test:
	$(GO) test .


clean:
	@echo "Cleaning SRP.."
	@rm -rvf main


proto: ready.proto
	@$(PROTOC) ready.proto --go_out=.
	@$(PROTOC) ready.proto --go_out=plugins=grpc:.


fmt:
	$(GO) fmt .

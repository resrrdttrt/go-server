PROTO_DIR := ./proto
GO_PB_DIR := ./user/pb

# Find all .proto files in the PROTO_DIR
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Default target
all: proto

# Generate Go and Py files from .proto files
proto: $(GO_PB_DIR) $(PROTO_FILES)
	rm -rf $(GO_PB_DIR)/*
	protoc --go_out=$(GO_PB_DIR) --go_opt=paths=source_relative --go-grpc_out=$(GO_PB_DIR) --go-grpc_opt=paths=source_relative -I $(PROTO_DIR) $(PROTO_FILES)

# Create PB_DIR if it doesn't exist
$(GO_PB_DIR):
	mkdir -p $(GO_PB_DIR)

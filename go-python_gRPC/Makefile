.PHONY: all proto clean run-server run-client

# Variables
PROTO_DIR = proto
PROTO_FILE = $(PROTO_DIR)/user.proto
SERVER_DIR = ./server
CLIENT_DIR = ./client
GOPATH = $(shell go env GOPATH)

all: proto

# Generate Proto files for both Go and Python
proto_go: 
	protoc --go_out=$(SERVER_DIR)/pkg \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(SERVER_DIR)/pkg \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILE)
proto_py: 
	# Generate Python files
	python3 -m grpc_tools.protoc \
		-I$(PROTO_DIR) \
		--python_out=$(CLIENT_DIR) \
		--grpc_python_out=$(CLIENT_DIR) \
		$(PROTO_FILE)


# Initialize Go modules
init-go:
	cd $(SERVER_DIR) && go mod init grpc-tutorial
	cd $(SERVER_DIR) && go mod tidy

# Install Python dependencies
init-python:
	pip3 install grpcio grpcio-tools

# Run the server
run-server:
	cd $(SERVER_DIR) && go run main.go

# Run the client
run-client:
	python3 $(CLIENT_DIR)/main.py

# Clean generated files
clean:
	rm -f $(CLIENT_DIR)/*pb2*.py
	rm -f $(SERVER_DIR)/pkg/user/*.go


# # Add protobuf repository
# curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v24.4/protoc-24.4-linux-x86_64.zip
# unzip protoc-24.4-linux-x86_64.zip -d $HOME/.local
# export PATH="$PATH:$HOME/.local/bin"

# # Or use apt (might be older version)
# sudo apt update
# sudo apt install -y protobuf-compiler
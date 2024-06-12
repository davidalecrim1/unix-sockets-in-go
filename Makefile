.PHONY: setup run-server run-client up

# Target to set up directory structure
setup:
	mkdir -p /tmp/go

# Target to run the server asynchronously
run-server: setup
	@echo "Starting server..."
	@GO_SOCKET_PATH=/tmp/go/unixsocket go run server/main.go &

# Target to run the client asynchronously
run-client: setup
	@echo "Starting client..."
	@GO_SOCKET_PATH=/tmp/go/unixsocket go run client/main.go &

# Target to run server and client sequentially
up: run-server run-client
	@echo "Waiting for server and client to finish..."
	@wait
	@echo "Server and client have finished."

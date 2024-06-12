# Unix Sockets in Go: Server-Client Communication

This project demonstrates how to establish communication between a server and client using Unix sockets in Go programming language. Unix sockets provide a mechanism for inter-process communication (IPC) within the same host, allowing efficient and reliable data exchange.

## What are Unix Sockets?

Unix sockets are a type of inter-process communication (IPC) mechanism provided by Unix-like operating systems. They allow communication between processes on the same host by creating a special file known as a socket file in the file system. Processes can connect to this socket and exchange data using standard networking APIs.

### Key Features of Unix Sockets:
- **Efficiency**: Data transfer occurs within the kernel, avoiding unnecessary copying.
- **Security**: Access to sockets can be controlled using file system permissions.
- **Local Communication**: Ideal for communication between processes running on the same machine.

## Getting Started

### Prerequisites

- Go 1.22 or higher

### Setting Up

1. **Clone the Repository**:

   ```bash
   git clone <repository-url>
   cd unix-sockets-go
   ```

2. **Setup the Directory Structure**:

   The Makefile included in this project automates the setup of necessary directories.

   ```bash
   make setup
   ```

   This command will create the directory `/tmp/go` where Unix socket files will be stored.

### Running the Server and Client

#### Starting the Server

To start the server, use the following command:

```bash
make run-server
```

This command compiles and runs the server program (`server/main.go`). The server will create a Unix socket at `/tmp/go/unixsocket` and listen for incoming connections.

#### Starting the Client

To start the client, use the following command in a separate terminal:

```bash
make run-client
```

This command compiles and runs the client program (`client/main.go`). The client will connect to the Unix socket created by the server (`/tmp/go/unixsocket`), send a message, and receive a response from the server.

### Running Both Server and Client Together

You can start both the server and client simultaneously using the following command:

```bash
make up
```

This command executes `make run-server` and `make run-client` in parallel, demonstrating concurrent communication between server and client via Unix sockets.
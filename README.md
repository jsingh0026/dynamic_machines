# Dynamic Machines

## Overview

**Dynamic Machines** is a simple, in-memory service that simulates machine creation, management, and cloning. It offers a basic set of functionalities such as starting, stopping, listing, and cloning machines through a gRPC API. This project serves as a working skeleton for simulating a machine orchestration system similar to Fly.io.

## Features

- **Start Machine**: Starts a new machine by specifying an ID and host.
- **Stop Machine**: Stops a machine by specifying its ID.
- **Clone Machine**: Clones an existing machine, creating a new machine with a different ID and host.
- **List Machines**: Retrieves a list of all machines, showing their current status.

## Tech Stack

- **Go**: Programming language used for the service.
- **gRPC**: Communication protocol for the service API.
- **Protocol Buffers (Protobuf)**: For defining the service API.
- **In-memory store**: Machines are stored in-memory for simplicity.

## Getting Started

### Prerequisites

1. **Go**: Ensure Go is installed on your machine. If not, install it from [golang.org](https://golang.org/dl/).
2. **Protocol Buffers**: Install `protoc` (Protocol Buffers compiler) to generate Go files from `.proto` files.

```bash
# Install protocol buffers compiler (protoc)
brew install protobuf # For macOS
# OR
# For Ubuntu
sudo apt install protobuf-compiler
```

3. **Install gRPC and Protobuf Go plugin**:

```bash
# Install gRPC for Go
go get google.golang.org/grpc

# Install Protobuf Go plugin for generating Go code from .proto files
go get google.golang.org/protobuf/cmd/protoc-gen-go
```

### Clone the repository

```bash
git clone https://github.com/your-repo/dynamic-machines.git
cd dynamic-machines
```

### Generate gRPC Go files

Run `protoc` to generate Go code from the `proto` definition:

```bash
# Assuming you have the proto file in the /proto directory
protoc --go_out=. --go-grpc_out=. proto/machines.proto
```

### Run the Service

1. Navigate to the root of the project and run the `main.go` file:

```bash
go run main.go
```

The server will start running on `localhost:50051`.

### Testing the Service

Once the service is running, you can interact with it via **gRPC** or **grpcurl**.

- **Start Machine**

```bash
grpcurl -plaintext -d '{"id": "machine1", "host": "host1.fly.dev"}' localhost:50051 proto.MachineService/StartMachine
```

- **Stop Machine**

```bash
grpcurl -plaintext -d '{"id": "machine1"}' localhost:50051 proto.MachineService/StopMachine
```

- **Clone Machine**

```bash
grpcurl -plaintext -d '{"oldId": "machine1", "newId": "machine2", "newHost": "host2.fly.dev"}' localhost:50051 proto.MachineService/CloneMachine
```

- **List Machines**

```bash
grpcurl -plaintext -d '{}' localhost:50051 proto.MachineService/ListMachines
```

### Sample Output

**Start Machine** response:

```json
{
  "id": "machine1",
  "host": "host1.fly.dev",
  "running": true,
  "started": "2025-04-29T20:11:07+05:30"
}
```

**Clone Machine** response:

```json
{
  "id": "machine2",
  "host": "host2.fly.dev",
  "running": true,
  "started": "2025-04-29T20:13:11+05:30"
}
```

## Project Structure

```
.
├── main.go                 # Main server and handlers
├── proto/                  # Protobuf definitions and generated Go files
│   └── machines.proto      # Protobuf definition file for MachineService
├── handler.go              # gRPC handler methods
├── README.md               # Project documentation
└── go.mod                  # Go module dependencies
```

## Next Steps

- **Persistence**: Implement disk persistence (e.g., using SQLite or Redis).
- **Health checks**: Add basic health checks for machines.
- **API expansion**: Add more features like scaling, status checks, and deployments.

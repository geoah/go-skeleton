# Sample gRPC Service

This project demonstrates a simple gRPC service using Go, Chi, Buf, and Connect-Go.

## Prerequisites

- [mise](https://mise.jdx.dev/) - A dev tool manager

## Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/geoah/go-skeleton.git
   cd go-skeleton
   ```

2. Set up mise:
   ```
   mise trust
   mise install
   ```

3. Run the service:
   ```
   air
   ```
   The service will start and listen on port 8080.

## Project Structure

- `proto/`: Contains the protobuf definitions
- `internal/`: Contains the service implementation
- `cmd/apiservice/`: Contains the main application code

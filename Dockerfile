FROM golang:1.23

WORKDIR /app

# Copy go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application (adjust if needed)
RUN go build -o main ./cmd/main

# Run the built binary
CMD ["./main"]
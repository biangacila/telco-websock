# Use an official Golang image as the base image
FROM golang:1.23.2-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files haven't changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app with optimizations for smaller binaries
RUN go build -o /programfile main.go

# Expose port 8080 to the outside world (this is where your service will run)
EXPOSE 8080

# Run the executable by default when the container starts
CMD ["/programfile"]

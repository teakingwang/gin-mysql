# Use the official Golang image to create a build artifact.
# This is based on Debian and has many common packages pre-installed.
# https://hub.docker.com/_/golang
FROM golang:1.18-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o my-gin-app .

# Use a lightweight base image to run the built binary
FROM alpine:latest AS runner
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/my-gin-app .

# Make the binary executable
RUN chmod +x my-gin-app

# Expose port 8080 to the Docker host
EXPOSE 8080

# Define environment variable
ENV NAME World

# Run the application
CMD ["./my-gin-app"]
# Build stage
FROM golang:alpine AS builder

# Install Git
RUN apk add --no-cache git

# Set working directory
WORKDIR /go/src/app

# Clone only necessary code from the repository
RUN git clone --depth 1 --branch main https://github.com/abohmeed/weatherapp.git

# Change working directory to the auth service
WORKDIR /go/src/app/weatherapp/auth

# Build the Go application and place the binary in /app
RUN go build -o /go/bin/app -v ./...

# Final stage
FROM alpine:latest

# Install necessary certificates
RUN apk --no-cache add ca-certificates

# Copy the built binary from the previous stage to /app
COPY --from=builder /app /app

# Set the entry point to the binary
ENTRYPOINT ["/app"]

# Expose port 8082 for communication
EXPOSE 8082

# Metadata indicating name and version of the service
LABEL Name=auth Version=0.1.0


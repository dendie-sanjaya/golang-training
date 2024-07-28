# Stage 1: Build stage
FROM golang:1.22.4-alpine AS build

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY app-tinyurl/go.mod app-tinyurl/go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY app-tinyurl/ ./

# Copy configuration file
COPY config-app.go /app/config/config.go

# Go mod tidy
RUN go mod tidy

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o go-tinyurl .

# Stage 2: Final stage
FROM alpine:edge

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/go-tinyurl .

# Set the entrypoint command
CMD ["./go-tinyurl"]
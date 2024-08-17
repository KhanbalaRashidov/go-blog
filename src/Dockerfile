# Start from the official Golang image
FROM golang:1.22.5

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY ./go.mod ./go.sum ./

# Download Go modules
RUN go mod download

# Copy the entire source directory into the container
COPY ./ ./

# Set the working directory to the directory containing the main.go file
WORKDIR /app/cmd/api

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /api

# Expose the application's port
EXPOSE 3003

# Command to run the API binary
CMD ["/api"]

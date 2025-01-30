# Use the official Golang image as the base image
FROM golang:1.22.6-alpine

# Set the working directory inside the container
WORKDIR /app_amit

# Copy the go.mod and go.sum files
COPY . .

# Download the Go module dependencies
RUN go mod download

# Expose the port the application runs on
EXPOSE 4001

CMD ["go", "run", "run.go"]

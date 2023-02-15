# Use an official Go image as the base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the source files to the container
COPY . .

# Download the dependencies
RUN go mod download

# Build the Go application
RUN go build -o go-task-mgr ./cmd

# Expose the port 9090
EXPOSE 9090

# Run the application when the container starts
CMD ["./go-task-mgr"]

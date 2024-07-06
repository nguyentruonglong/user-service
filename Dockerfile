# Use the official Golang image as a base image
FROM golang:1.21.1

# Install the necessary build tools
RUN apt-get update && apt-get install -y gcc musl-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o user-service main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable with the production configuration
CMD ["./user-service", "--config=config/dev_config.yaml"]

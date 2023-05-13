FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
RUN mkdir ./app
WORKDIR ./app

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./

# Build the Go app
RUN go build -o main ./cmd

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the executable
CMD ["./main"]
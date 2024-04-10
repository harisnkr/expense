# Use the official golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Fetch the dependencies
RUN go mod download

# Build the Go binary
RUN go build -o main .

# Expose port 8801 to the outside world
EXPOSE 8801

# Command to run the executable
CMD ["./main"]

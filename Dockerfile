# Specify a base image for the Dockerfile
FROM golang:1.18 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files for dependency management
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application by compiling it
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start the next stage to create a smaller base image
FROM alpine:latest as runner

# Add certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Set the working directory in the new stage
WORKDIR /root/

# Copy the compiled binary from the Builder stage to the Runner stage
COPY --from=builder /app/main .

# The critical step: Copy the HTML templates and static files to the correct location
COPY --from=builder /app/html /root/html

# Define the command to run the application
CMD ["./main"]

# Instructions for building and running the Docker Image (these comments are for reference and should be run in the terminal):
# 1. Build the Docker Image
#    docker build -t my-go-app .
# 2. Run the Docker Container:
#    docker run -p 8080:8080 my-go-app

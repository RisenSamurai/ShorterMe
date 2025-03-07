# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /backend

# Copy dependency files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your source code
COPY . .

# Build the Go app into a binary named "main"
RUN go build -o main .

# Expose the port that the app listens on (e.g., 8080)
EXPOSE 8080

# Command to run the binary
CMD ["./main"]

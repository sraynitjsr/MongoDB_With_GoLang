# Use the official Golang image as the base image
FROM golang:1.16-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8000 for the application
EXPOSE 8000

# Command to run the executable
CMD ["./main"]

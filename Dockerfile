# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files to the container
COPY . .

# Download Go dependencies
RUN go get -d -v ./...

# Build the Go application
RUN go build -o main .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]

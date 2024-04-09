# Use the official Golang image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Download Go dependencies
RUN go mod download

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the application
CMD ["air"]

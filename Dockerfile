# Use the official Golang image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.0/migrate.linux-amd64.tar.gz | tar xvz && mv ./migrate /usr/local/bin

RUN go install github.com/golang/mock/mockgen

# Download Go dependencies
RUN go mod download

# Copy the entrypoint script into the container
COPY docker-entrypoint.sh /entrypoint.sh

# Set the entrypoint
ENTRYPOINT ["/entrypoint.sh"]

# Expose port 8181 to the outside world
EXPOSE 8181

BINARY_NAME=quake-data-collector.out
 
MAIN_FILE="internal/context/log_handler/main.go"

build:
	@go build -ldflags "-s -w" -o ${BINARY_NAME} ${MAIN_FILE}
 
test: .FORCE
	@go test -v ./...

run:
	@go build -ldflags "-s -w" -o ${BINARY_NAME} ${MAIN_FILE}
	@./${BINARY_NAME}
 
clean:
	@go clean
	@rm ${BINARY_NAME}

.FORCE:
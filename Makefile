BINARY_NAME=go-ddd.out
 
include .env
export $(shell sed 's/=.*//' .env)

MAIN_FILE="cmd/log_handler/main.go"
DATABASE_URL=mysql://root:secret@127.0.0.1:3306/db?sslmode=disable
MIGRATIONS_LOG_HANDLER_PATH=./migrations/log_handler
DATABASE_URL=mysql://$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_DATABASE)

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

migrate-up:
	migrate -path $(MIGRATIONS_LOG_HANDLER_PATH) -database "$(DATABASE_URL)" -verbose up

migrate-down:
	migrate -path $(MIGRATIONS_LOG_HANDLER_PATH) -database "$(DATABASE_URL)" -verbose down

.FORCE:
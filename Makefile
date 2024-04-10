BINARY_NAME=go-ddd.out
 
include .env
export $(shell sed 's/=.*//' .env)

MAIN_FILE="cmd/log_handler/main.go"
MIGRATIONS_LOG_HANDLER_PATH=./db/migrations/log_handler
DATABASE_URL=mysql://$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_DATABASE)

local.build:
	@go build -ldflags "-s -w" -o ${BINARY_NAME} ${MAIN_FILE}
 
local.test: .FORCE
	@go test -v ./...

local.bench: .FORCE
	@go test -v ./... -bench=.	

local.run:
	@go build -ldflags "-s -w" -o ${BINARY_NAME} ${MAIN_FILE}
	@./${BINARY_NAME}
 
local.clean:
	@go clean
	@rm ${BINARY_NAME}

local.migrate-up:
	migrate -path $(MIGRATIONS_LOG_HANDLER_PATH) -database "$(DATABASE_URL)" -verbose up

local.migrate-down:
	migrate -path $(MIGRATIONS_LOG_HANDLER_PATH) -database "$(DATABASE_URL)" -verbose down

local.debug: .FORCE
	cd cmd/log_handler && dlv debug --headless --listen=:2345 --api-version=2

docker-compose.run:
	docker-compose exec app go run ${MAIN_FILE}

.FORCE:
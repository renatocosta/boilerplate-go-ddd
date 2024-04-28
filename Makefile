BINARY_NAME=go-ddd
 
include .env
export $(shell sed 's/=.*//' .env)

CMD_LOG_HANDLER=cmd/log_handler
START_MAIN_FILE=$(CMD_LOG_HANDLER)/main.go
START_LOG_HANDLER_WORKER_WORKFLOW="$(CMD_LOG_HANDLER)/worker_workflow/main.go"
START_MATCH_REPORTING_WORKER_WORKFLOW="cmd/match_reporting/worker_workflow/main.go"
DATABASE_URL=mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)
 
run: .FORCE
	go build -ldflags "-s" -v -o ./tmp/${BINARY_NAME} ./${CMD_LOG_HANDLER}

## Start Temporal Workers Workflow 
run.log_handler_worker_workflow: .FORCE
	go run ${START_LOG_HANDLER_WORKER_WORKFLOW}

run.match_reporting_worker_workflow: .FORCE
	go run ${START_MATCH_REPORTING_WORKER_WORKFLOW}

race: .FORCE
	go run -race github.com/ddd/cmd/log_handler

test: .FORCE
	go test -v ./...

test.race: .FORCE
	go test -race ./...

bench: .FORCE
	go test -v ./... -bench=.

migrate-up:
	migrate -path $(CMD_LOG_HANDLER)/migrate -database "$(DATABASE_URL)" -verbose up

migrate-down:
	migrate -path $(CMD_LOG_HANDLER)/migrate -database "$(DATABASE_URL)" -verbose down

seed:
	go run $(CMD_LOG_HANDLER)/seed/main.go

clean:
	go clean
	rm ${BINARY_NAME}

mockgen:
	mockgen -source=$(source) -destination=$(destination) -package=$(package)

debug: .FORCE
	cd cmd/log_handler && dlv debug --headless --listen=:2345 --api-version=2
	
.FORCE:
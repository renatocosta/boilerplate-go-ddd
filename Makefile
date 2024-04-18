BINARY_NAME=go-ddd.out
 
include .env
export $(shell sed 's/=.*//' .env)

START_MAIN_FILE="cmd/log_handler/main.go"
START_LOG_HANDLER_WORKER_WORKFLOW="cmd/log_handler/worker_workflow/main.go"
START_MATCH_REPORTING_WORKER_WORKFLOW="cmd/match_reporting/worker_workflow/main.go"
MIGRATIONS_LOG_HANDLER_PATH=./db/migrations/log_handler
DATABASE_URL=mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

run: .FORCE
	cd deployments && docker-compose exec app go run ${START_MAIN_FILE}	
 
## Start Workers Workflow 
run.log_handler_worker_workflow: .FORCE
	cd deployments && docker-compose exec app go run ${START_LOG_HANDLER_WORKER_WORKFLOW}

run.match_reporting_worker_workflow: .FORCE
	cd deployments && docker-compose exec app go run ${START_MATCH_REPORTING_WORKER_WORKFLOW}

test: .FORCE
	cd deployments && docker-compose exec app go test -v ./...

bench: .FORCE
	cd deployments && docker-compose exec app go test -v ./... -bench=.

migrate-up:
	cd deployments && docker-compose exec app migrate -path $(MIGRATIONS_LOG_HANDLER_PATH) -database "$(DATABASE_URL)" -verbose up

migrate-down:
	cd deployments && docker-compose exec app migrate -path $(MIGRATIONS_LOG_HANDLER_PATH) -database "$(DATABASE_URL)" -verbose down

clean:
	cd deployments && docker-compose exec app go clean
	cd deployments && docker-compose exec app rm ${BINARY_NAME}

mockgen:
	cd deployments && docker-compose exec app mockgen -source=$(source) -destination=$(destination) -package=$(package)

debug: .FORCE
##	cd cmd/log_handler && dlv debug --headless --listen=:2345 --api-version=2
    cd deployments && docker-compose exec app cd cmd/log_handler && dlv debug --headless --listen=:2345 --api-version=2
.FORCE:
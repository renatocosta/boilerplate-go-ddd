#!/bin/bash

# Run the log_handler main.go
go run cmd/log_handler/main.go &

# Run the log_handler worker_workflow main.go
go run cmd/log_handler/worker_workflow/main.go &

# Run the match_reporting worker_workflow main.go
go run cmd/match_reporting/worker_workflow/main.go &

make migrate up

# Start your application
exec air

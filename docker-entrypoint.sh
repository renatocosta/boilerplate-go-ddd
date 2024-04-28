#!/bin/bash

# Run the log_handler main.go
make run &

# Run the log_handler worker_workflow main.go
make run.log_handler_worker_workflow &

# Run the match_reporting worker_workflow main.go
make run.match_reporting_worker_workflow &

make migrate-up

# Start your application
exec air

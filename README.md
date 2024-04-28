# Strategic Monolith ⚡️

Example of Clean Architecture and Domain Driven Design + Workflow Engine(Temporal.io) in Golang

## Unit testing
```
make test
```

## Let's Run the Application 
```
docker-compose up
or
make migrate-up
make run
make run.log_handler_worker_workflow
make run.match_reporting_worker_workflow
```

## Event Modelling

Go through all of the learning journey using Event Modelling for understanding the business needs as shown below

### Steps
![Image](./assets/EventModelling.jpg?raw=true)

## Bounded contexts
![Image](./assets/EventModellingOutcome.jpg?raw=true)
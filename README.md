# agileEngine Golang simple Test

This project was builded with go ver 1.12.7.
also is use almost all default library on Golang only used external UUID for inserts
its use RWMutex for concurrency locking,
by default all call on http service is executed on gorutine

## Development server

Run `go run main.go` for a dev server. Navigate to `http://localhost:8001/`.

## Routes
post transaction
http://localhost:8001/transaction

get transaction by Id
http://localhost:8001/transaction/:Id

get array of transactions
http://localhost:8001/alltransactions

get balance
http://localhost:8001/


## Build

Run `go build -ldflags="-s -w"` to build the project withOut debugin. 

## UX/UI
go to angular/agile-engine-ui and see the README.md

## Further help

To get more help on the project just email juanvalero@gmail.com.

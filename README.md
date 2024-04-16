# Microservice-PB
This sample code is submitted in partial fulfilment of the requirement by ENU for MSC in Computing

## Objectives
- Test the hypothesis of priority-based queue on asynchronous communication in microservices
- Evaluating the performance of priority-based queue with ordered queue in asynchronous communication in microservices

## Configuration References
- https://kafka.apache.org/documentation

 
## Environment Setup
- Install Go version 1.21.5 (https://go.dev/doc/devel/release#go1.21)
- run `go version` to make sure go is available on the terminal

## Starting the project
- run command `go tidy`  (This will install all the necessary packages)
- cd into the the main.go file and run `go run main.go` and see the terminal for any notification

## Navigating on the repo
At the time of this Readme, we have about 6 branches of test case on this hypothesis, for this submission we would only be reviewing **2 branches** as follows;

- `ft-test` (the test of the hypothesis )
- `multi-producer-test` ( The performance testing between priority-based and ordered queues)

## Packages
### Installed packages
Golang client for Apache Kafka REF: https://github.com/confluentinc/confluent-kafka-go

-  ` go get -u github.com/confluentinc/confluent-kafka-go/v2/kafka `

Viper for loading configuration data REF: https://github.com/spf13/viper

-  ` go get github.com/spf13/viper`

ORM for Golang REF: https://gorm.io/
- ` go get -u gorm.io/gorm `

GORM driver for Sqlite REF: https://github.com/go-gorm/sqlite 
 - `go get -u gorm.io/driver/sqlite `









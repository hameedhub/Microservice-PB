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
- create and start the docker container using the following commands on the docker file directory `docker-compose up` and `docker-compose down` to stop the container
- cd into the the main.go file and run `go run main.go` for `ft-test` and see the terminal for any notification
  
- For `mult-producer-test`  run `sh run.sh` to run the producers and `sh run_consmer.sh` to start the consumer.

## Navigating on the repo
At the time of this Readme, we have about 6 branches of test case on this hypothesis, for this submission we would only be reviewing **2 branches** as follows;

- `ft-test` (the test of the hypothesis )
- `multi-producer-test` ( The performance testing between priority-based and ordered queues)

## Installed packages
Golang client for Apache Kafka REF: https://github.com/confluentinc/confluent-kafka-go

-  ` go get -u github.com/confluentinc/confluent-kafka-go/v2/kafka `

Viper for loading configuration data REF: https://github.com/spf13/viper

-  ` go get github.com/spf13/viper`

ORM for Golang REF: https://gorm.io/
- ` go get -u gorm.io/gorm `

GORM driver for Sqlite REF: https://github.com/go-gorm/sqlite 
 - `go get -u gorm.io/driver/sqlite `

## Results
![image](https://github.com/hameedhub/Microservice-PB/assets/46590803/02d1734d-3bdb-4226-a365-305fb6edc715)
![image](https://github.com/hameedhub/Microservice-PB/assets/46590803/18085a29-4285-4a8f-9055-661e864aa44b)
![image](https://github.com/hameedhub/Microservice-PB/assets/46590803/06440740-fd87-41bb-a7c3-68eafeb05f13)
![image](https://github.com/hameedhub/Microservice-PB/assets/46590803/2fcfa1c7-4083-4ca9-a227-47c88ca41f35)
![image](https://github.com/hameedhub/Microservice-PB/assets/46590803/df65340a-b964-4608-b90e-00d113be64bd)











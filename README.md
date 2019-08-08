# kafkapoc
Exploring the innards of Sarama library for Kafka. Language used is Golang

#Usage:

To run Kafka producer
```text
go run main.go producer --topic TOPICNAME --message MESSAGESTRING
```

Note : List of topics is specified in the config file. The topic name used should be from the list.

To run Kafka consumer
```text
go run main.go consumer
```
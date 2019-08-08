package services

import (
	cluster "github.com/bsm/sarama-cluster"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func Consume(brokers []string, topics []string) {
	logrus.Info("Inside consumer function")
	config := cluster.NewConfig()
	//topic := topics[0]

	// SARAMA-CLUSTER CONSUMER IMPLEMENTATION

	consumer, err := cluster.NewConsumer(brokers, "cgroup1", topics, config)
	if err != nil {
		logrus.WithError(err).Error("Error creating a consumer instance")
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)


	// Messages are multiplexed and made available to the consuming party. Using default mode for this.
	// else will have to listen for messages in each individual partition
	logrus.Info("Created consumer.")
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				logrus.WithFields(logrus.Fields{
					"Partition": msg.Partition,
					"Offset": msg.Offset,
					"Value": string(msg.Value),
				}).Info("Printing data")
				consumer.MarkOffset(msg, "")	// mark message as processed
			}
		case <-signals:
			return
		}
	}

	// SARAMA CONSUMER IMPLEMENTATION. Uncomment to include
	//config := sarama.NewConfig()
	//consumer, err := sarama.NewConsumer(brokers, config)
	//if err != nil {
	//	logrus.WithError(err).Error("Error creating consumer")
	//}
	//partitionList, _ := consumer.Partitions(topic)
	//offset := sarama.OffsetOldest
	//for partition := range partitionList {
	//	pc, _ := consumer.ConsumePartition(topic, int32(partition), offset)
	//	for message := range pc.Messages() {
	//		logrus.WithFields(logrus.Fields{
	//			"Key": message.Key,
	//			"Value": string(message.Value),
	//		}).Info("Message Received")
	//	}
	//}
}

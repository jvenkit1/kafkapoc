package services

import(
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.Info("Inside init function")
}

func Produce(brokers []string, consumerMessage string, topicName string) {

	config := sarama.NewConfig()
	config.Producer.Partitioner=sarama.NewRandomPartitioner
	config.Producer.RequiredAcks=sarama.WaitForAll
	config.Producer.Return.Successes=true

	// send message with sync producer first, learn about channels and then send it with async producer
	//producer, err := sarama.NewAsyncProducer(brokers, config)
	//if err != nil {
	//	logrus.WithError(err).Error("Failed generating a producer instance")
	//}

	p2, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		logrus.WithError(err).Error("Failed generating a producer instance")
	}
	msg := &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.StringEncoder(consumerMessage),
	}

	partition, offset, err := p2.SendMessage(msg)
	if err != nil {
		logrus.WithError(err).Error("Failed to send a message")
	}

	logrus.WithFields(logrus.Fields{
		"Partition": partition,
		"Offset": offset,
		"Topic": msg.Topic,
		"Message": consumerMessage,
	}).Info("Sent message")
}

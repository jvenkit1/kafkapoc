package main

import (
	"github.com/sirupsen/logrus"
	"kafkapoc/services"
)

func main() {
	logrus.Info("Hello World")
	services.Produce()
	logrus.Info("Moving to consumer now")
	services.Consume()
}

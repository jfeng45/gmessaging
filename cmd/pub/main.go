package main

import (
	"github.com/jfeng45/gmessaging"
	"github.com/jfeng45/gmessaging/config"
	"github.com/jfeng45/gmessaging/factory"
	"github.com/nats-io/nats.go"
	"log"
)

type PaymentEvent struct {
	Id int
	SourceAccount string
}
func main() {
	subject :="test"
	p := PaymentEvent{1, "sourceAccount"}
	mi, err := initMessagingService()
	if err != nil {
		log.Println("err:", err)
	}
	err = mi.Publish(subject, p)
	if err != nil {
		log.Println("err:", err)
	}
	err = mi.Flush()
	if err != nil {
		log.Println("err:", err)
	}
	mi.Close()
	mi.CloseConnection()
}

func initMessagingService() (gmessaging.MessagingInterface, error) {
	config := config.Messaging{config.NATS_ENCODED, nats.DefaultURL, nats.JSON_ENCODER}
	return factory.Build(&config)
}


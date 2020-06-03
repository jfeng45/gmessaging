package main

import (
	"github.com/jfeng45/gmessaging"
	 "github.com/jfeng45/gmessaging/config"
	 "github.com/jfeng45/gmessaging/factory"
	 "github.com/nats-io/nats.go"
	 "log"
	 "runtime"
)
type PaymentEvent struct {
	Id int
	SourceAccount string
}

func main() {
	subject :="test"
	mi, err := InitMessagingService()
	if err != nil {
		log.Println("err:", err)
	}
	_, err = mi.Subscribe(subject, func(pe PaymentEvent) {
		log.Println("payload:",pe)
	})
	if err != nil {
		log.Println("err:",err)
	}
	log.Printf("Listening on [%s]", subject)
	runtime.Goexit()
}

func InitMessagingService() (gmessaging.MessagingInterface, error) {
	config := config.Messaging{config.NATS_ENCODED, nats.DefaultURL, nats.JSON_ENCODER}
	return factory.Build(&config)
}

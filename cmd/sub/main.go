package main

import (
	"github.com/jfeng45/gmessaging"
	"github.com/jfeng45/gmessaging/nat"
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
	mi, err := initMessagingService()
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

func initMessagingService() (gmessaging.MessagingInterface, error) {
	url := nats.DefaultURL
	nc, err :=nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	//defer nc.Close()
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	//defer ec.Close()
	if err != nil {
		return nil, err
	}
	n := nat.Nat{ec}
	return &n, nil

}

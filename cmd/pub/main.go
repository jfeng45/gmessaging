package main


import (
	"github.com/jfeng45/gmessaging"
	"github.com/jfeng45/gmessaging/nat"
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
	return n, nil
}

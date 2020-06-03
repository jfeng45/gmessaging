package factory

import (
	"errors"
	"github.com/jfeng45/gmessaging"
	"github.com/jfeng45/gmessaging/config"
	"github.com/jfeng45/gmessaging/nat"
	"github.com/nats-io/nats.go"
	"log"
)

func Build(c *config.Messaging) (gmessaging.MessagingInterface, error) {
	switch c.Code {
	case config.NATS_ENCODED:
		return buildNatsEncoded(c)
	default:
		msg := "No such code '" + c.Code + "' for messging service"
		return nil, errors.New(msg)
	}

}

func buildNatsEncoded(c *config.Messaging) (gmessaging.MessagingInterface, error){
	url := c.ServerUrl
	nc, err :=nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	//defer nc.Close()
	ec, err := nats.NewEncodedConn(nc, c.Encoder)
	if err != nil {
		return nil, err
	}
	nat := nat.Nat{ec}
	return &nat, nil
}

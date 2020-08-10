package nat

import (
	"github.com/jfeng45/gmessaging"
	"github.com/jfeng45/gmessaging/config"
	"github.com/jfeng45/gmessaging/nat/natsencoded"
	"github.com/jfeng45/gmessaging/nat/natsnonencoded"
	"github.com/nats-io/nats.go"
	"log"
)

func BuildNatsEncoded(c *config.Messaging) (gmessaging.MessagingEncodedInterface, error){
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
	nat := natsencoded.NatEncoded{ec}
	return &nat, nil
}

func BuildNatsNonEncoded(c *config.Messaging) (gmessaging.MessagingInterface, error){
	url := c.ServerUrl
	nc, err :=nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	////defer nc.Close()
	//ec, err := nats.NewEncodedConn(nc, c.Encoder)
	//if err != nil {
	//	return nil, err
	//}
	nne := natsnonencoded.Nat{nc}
	return &nne, nil
}

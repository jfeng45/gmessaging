package natsencoded

import (
	"github.com/nats-io/nats.go"
)

type NatEncoded struct {
	Ec *nats.EncodedConn
}

func (n *NatEncoded) Publish(subject string, i interface{}) error {
	return n.Ec.Publish(subject,i)
}

func (n *NatEncoded) Subscribe(subject string, i interface{} ) (interface{}, error) {
	h := i.(nats.Handler)
	subscription, err :=n.Ec.Subscribe(subject, h)
	return subscription, err
}

func  (n *NatEncoded) Flush() error {
	return n.Ec.Flush()
}

func  (n *NatEncoded) Close()  {
	n.Ec.Close()
}

func  (n *NatEncoded) CloseConnection()  {
	n.Ec.Conn.Close()
}
package nat

import (
	"github.com/nats-io/nats.go"
)

type Nat struct {
	Ec *nats.EncodedConn
}

func (n *Nat) Publish(subject string, i interface{}) error {
	return n.Ec.Publish(subject,i)
}

func (n *Nat) Subscribe(subject string, i interface{} ) (interface{}, error) {
	h := i.(nats.Handler)
	subscription, err :=n.Ec.Subscribe(subject, h)
	return subscription, err
}

func  (n *Nat) Flush() error {
	return n.Ec.Flush()
}

func  (n *Nat) Close()  {
	n.Ec.Close()
}

func  (n *Nat) CloseConnection()  {
	n.Ec.Conn.Close()
}
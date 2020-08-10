package natsnonencoded

import (
	"github.com/jfeng45/gmessaging"
	"github.com/nats-io/nats.go"
)

type Nat struct {
	Conn *nats.Conn
}

func (n *Nat) Publish(subject string, data []byte) error {
	return n.Conn.Publish(subject,data)
}

func (n *Nat) Subscribe(subject string,  mh gmessaging.GMessageHandler) (*gmessaging.GSubscription, error) {
	mhNats := func (msg *gmessaging.GMessage) func(*nats.Msg){
		a := func(nm *nats.Msg) {
			mh(toGMsg(nm))
		}
		return a
	}
	var msg *gmessaging.GMessage
	subscription, err :=n.Conn.Subscribe(subject, mhNats(msg))
	gs := toGSubscription(subscription)
	return gs, err
}

func  (n *Nat) Flush() error {
	return n.Conn.Flush()
}

func  (n *Nat) Close()  {
	n.Conn.Close()
}

func toGMsg(ns *nats.Msg)  *gmessaging.GMessage{
	if ns != nil {
		return &gmessaging.GMessage{Subject:ns.Subject, Reply:ns.Reply, Data:ns.Data }
	}
	return nil
}

func toGSubscription(s  *nats.Subscription) *gmessaging.GSubscription{
	if s != nil {
		return &gmessaging.GSubscription{Subject:s.Subject }
	}
	return nil
}

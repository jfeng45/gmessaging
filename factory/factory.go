package factory

import (
	"errors"
	"github.com/jfeng45/gmessaging"
	"github.com/jfeng45/gmessaging/config"
	"github.com/jfeng45/gmessaging/factory/nat"
)

func Build(c *config.Messaging) (gmessaging.MessagingInterface, error) {
	switch c.Code {
	case config.CODE_NATS:
		return nat.BuildNatsNonEncoded(c)
	default:
		msg := "No such code '" + c.Code + "' for messaging service"
		return nil, errors.New(msg)
	}
}

func BuildEncoded(c *config.Messaging) (gmessaging.MessagingEncodedInterface, error) {
	switch c.Code {
	case config.CODE_NATS:
		return nat.BuildNatsEncoded(c)
	default:
		msg := "No such code '" + c.Code + "' for Encoded messaging service"
		return nil, errors.New(msg)
	}
}

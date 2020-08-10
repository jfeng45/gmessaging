package cmd

import (
	"github.com/jfeng45/gmessaging"
	"github.com/jfeng45/gmessaging/config"
	"github.com/jfeng45/gmessaging/factory"
	"github.com/nats-io/nats.go"
)

func InitMessagingEncodedService() (gmessaging.MessagingEncodedInterface, error) {
	config := config.Messaging{config.CODE_NATS, nats.DefaultURL, nats.JSON_ENCODER}
	return factory.BuildEncoded(&config)
}

func InitMessagingService() (gmessaging.MessagingInterface, error) {
	// Non-encoded connection doesn't need encoder
	encoder := ""
	config := config.Messaging{config.CODE_NATS, nats.DefaultURL, encoder}
	return factory.Build(&config)
}


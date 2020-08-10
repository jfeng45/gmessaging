package gmessaging

// GMessageHandler is a callback function that processes messages delivered to asynchronous subscribers.
// It is a generic version of NATS's MsgHandler
type GMessageHandler func(msg *GMessage)

// GMessage is a structure used by Subscribers and PublishMsg().
// It is a generic version of NATS's Msg
type GMessage struct {
	Subject string
	Reply   string
	Data    []byte
	Sub     *GSubscription
}

// A GSubscription represents interest in a given subject.
// I am not sure how to use it, so only put one attribute in it
// It is a generic version of NATS's Subscription
type GSubscription struct {
	Subject string
}

type MessagingEncodedInterface interface {
	Publish(subject string, i interface{}) error
	Subscribe(subject string, cb interface{} ) (interface{}, error)
	Flush() error
	// Close will close the decorated connection (For example, it could be the coded connection)
	Close()
	// CloseConnection will close the connection to the messaging server. If the connection is not decorated, then it is
	// the same with Close(), otherwise, it is different
	CloseConnection()
}

type MessagingInterface interface {
	Publish(subject string, data []byte) error
	Subscribe(subject string, mh GMessageHandler) (*GSubscription, error)
	Flush() error
	// Close will close the decorated connection (For example, it could be the coded connection)
	Close()
}
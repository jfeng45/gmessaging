package gmessaging

type MessagingInterface interface {
	Publish(subject string, i interface{}) error
	Subscribe(subject string, cb interface{} ) (interface{}, error)
	Flush() error
	// Close will close the decorated connection (For example, it could be the coded connection)
	Close()
	// CloseConnection will close the connection to the messaging server. If the connection is not decorated, then it is
	// the same with Close(), otherwise, they are different
	CloseConnection()
}
package config

type Messaging struct {
	// code for the implementation messaging lib
	Code string `yaml:"code"`
	// Server URL for the messaging server
	ServerUrl string `yaml:"serverUrl"`
	// Encoder will be used by the lib
	Encoder string `yaml:"encoder"`
}

const (
	CODE_NATS         string = "codeNats"
)


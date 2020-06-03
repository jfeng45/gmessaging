package config

type Messaging struct {
	// log library name
	Code string `yaml:"code"`
	// log level
	ServerUrl string `yaml:"serverUrl"`
	// show caller in log message
	Encoder string `yaml:"encoder"`
}

const (
	NATS_ENCODED         string = "natsEncoded"
)


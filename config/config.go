package config

type Config struct {
	QueueSize int `envconfig:"QUEUE_SIZE"`
}

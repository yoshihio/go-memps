package entity

type Data struct {
	Topic   string
	Message string
}

func NewData(topic, message string) Data {
	return Data{
		Topic:   topic,
		Message: message,
	}
}

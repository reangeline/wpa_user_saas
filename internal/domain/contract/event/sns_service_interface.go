package contract

type SNSServiceInterface interface {
	PublishMessage(message string, subject string) error
	SubscribeToTopic(endpoint string, protocol string) (string, error)
}

package implementation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNSService struct {
	snsClient *sns.SNS
	topicArn  string
}

func NewSNSService(topicArn string) *SNSService {
	sess := session.Must(session.NewSession())
	return &SNSService{
		snsClient: sns.New(sess),
		topicArn:  topicArn,
	}
}

func (s *SNSService) PublishMessage(message string, subject string) error {
	input := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(s.topicArn),
		Subject:  aws.String(subject),
	}

	_, err := s.snsClient.Publish(input)
	return err
}

func (s *SNSService) SubscribeToTopic(endpoint string, protocol string) (string, error) {

	input := &sns.SubscribeInput{
		Endpoint: aws.String(endpoint),
		Protocol: aws.String(protocol),
		TopicArn: aws.String(s.topicArn),
	}

	result, err := s.snsClient.Subscribe(input)
	if err != nil {
		return "", err
	}
	return *result.SubscriptionArn, nil
}

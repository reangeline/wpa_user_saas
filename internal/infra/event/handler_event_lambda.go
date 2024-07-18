package event

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/reangeline/wpa_user_saas/internal/di"
)

type SNSMessage struct {
	Message   string `json:"Message"`
	Timestamp string `json:"Timestamp"`
}

type HandlerEventLambda struct {
	svc *dynamodb.DynamoDB
}

func NewHandlerEventLambda(svc *dynamodb.DynamoDB) *HandlerEventLambda {
	return &HandlerEventLambda{
		svc,
	}
}

func (hel *HandlerEventLambda) Handler(snsEvent events.SNSEvent) {

	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		switch snsRecord.Subject {
		case "created_user":

			user, err := di.InitializeEventUser(hel.svc)
			if err != nil {
				fmt.Println(err)
			}

			user.CreateUserLambda(snsRecord)
		default:
			fmt.Println("No subject found")
		}

	}

}
